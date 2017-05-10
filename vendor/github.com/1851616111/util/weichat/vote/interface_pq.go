package vote

import (
	"fmt"
	"sort"
	"time"
	"errors"

	"database/sql"
	"github.com/golang/glog"
	"github.com/lib/pq"
	"github.com/1851616111/util/weichat/manager/user"
	"github.com/1851616111/util/weichat/event"
)

var TABLE_VOTER string = "GO_WEICHAT_ACTIVITY_VOTER"
var SQL_TABLE = `
CREATE TABLE IF NOT EXISTS  GO_WEICHAT_ACTIVITY_VOTER (
   openid varchar(50) PRIMARY KEY,
   voterid SERIAL,
   name varchar(15),
   image varchar(300),
   company varchar(100),
   mobile varchar(20),
   declaration varchar(50),
   votedcount integer DEFAULT 0,
   voteRecords integer[],

   followed BOOLEAN DEFAULT FALSE,
   registed BOOLEAN DEFAULT FALSE,
   imageCached	BOOLEAN DEFAULT FALSE
);

CREATE OR REPLACE FUNCTION wc_Regist_Voter(p_openid varchar, p_name VARCHAR, p_image VARCHAR, p_company VARCHAR, p_mobile VARCHAR, p_declaration varchar, p_imageCached boolean) RETURNS VARCHAR AS $$
  DECLARE
    b_followed BOOLEAN = NULL;
    b_registed BOOLEAN = FALSE;
  BEGIN
    SELECT followed from go_weichat_activity_voter where openid = p_openid into b_followed;
    IF b_followed IS NULL THEN
         RETURN 'not found';
    ELSEIF b_followed = FALSE THEN
         RETURN 'not follow';
    ELSE
        SELECT registed from go_weichat_activity_voter where openid = p_openid into b_registed;
        IF b_registed THEN
              RETURN 'already regist';
        ELSE
              UPDATE go_weichat_activity_voter SET
                name=p_name, image=p_image, company=p_company, mobile=p_mobile, declaration=p_declaration, registed=TRUE , imageCached=p_imageCached
              WHERE openid = p_openid;
              RETURN 'ok';
        END IF;
    END IF;
  END;
$$ LANGUAGE plpgsql;`

var ERR_NOT_FOLLOW error = errors.New("user not follow")
var ERR_NOT_REGISTER error = errors.New("user not regist")
var ERR_EXIST error =  errors.New("user exists")
var ERR_UNKNOWN_STATUS error = errors.New("user not follow but regist")

func NewDBInterface(ip, port, user, passwd, database string) (DBInterface, error) {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, database)

	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

type DB struct {
	*sql.DB
}

func (d DB) Init(access_token string) error {
	if _, err := d.Exec(SQL_TABLE); err != nil {
		return err
	}

	users, err := user.ListUserIDs(access_token)
	if err != nil {
		return err
	}

	var cuccessCount int
	for id := range users {
		if _, err:= d.Exec(`INSERT INTO ` +TABLE_VOTER+ `(openid, followed) VALUES($1, TRUE) ON CONFLICT(openid)
		DO UPDATE SET openid = EXCLUDED.openid, followed = EXCLUDED.followed`, users[id]); err != nil {
			glog.Errorf("weichat init vote user err %v\n", err)
		} else {
			cuccessCount ++
		}
	}

	glog.Errorf("weichat activity init followers: %d \n", len(users))
	glog.Errorf("weichat activity init db user: %d \n", cuccessCount )
	return nil
}

func (d DB) Register(v *Voter) error {
	var result string
	if err := d.QueryRow(`SELECT wc_Regist_Voter($1, $2, $3, $4, $5, $6, $7)`,
		v.OpenID, v.Name, v.Image, v.Company, v.Mobile, v.Declaration, v.imageCached).Scan(&result); err != nil {
		fmt.Println(err)
	}

	switch result {
	case "already regist":
		return ERR_EXIST
	case "not found":
		return ERR_UNKNOWN_STATUS
	case "not follow":
		return ERR_NOT_FOLLOW
	case "ok":
		return nil
	default:
		return fmt.Errorf("unknown err %s\n", result)
	}
}

func (d DB) Vote(openID, votedID string) (err error) {
	voteRecords := pq.Int64Array{}
	if err = d.QueryRow(`SELECT voteRecords FROM `+TABLE_VOTER+` WHERE openid = $1`, openID).Scan(&voteRecords); err != nil {
		if err == sql.ErrNoRows {
			if _, err := d.Exec(`INSERT INTO ` +TABLE_VOTER+ ` (openid) VALUES($1)`, openID); err != nil {
				glog.Errorf("weichat vote from openid %s update user err %v\n", err)
				return err
			}
		} else {
			return
		}
	}

	records := []int64(voteRecords)
	if !hasVoteRight(records) {
		err = errors.New("vote chance large than 3")
		return
	}

	var tx *sql.Tx
	tx, err = d.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec(`UPDATE `+TABLE_VOTER+` SET votedcount = votedcount + 1 where voterid = $1`, votedID)
	if err != nil {
		return
	}

	voteRecords = append(voteRecords, time.Now().Unix())
	_, err = tx.Exec(`UPDATE `+TABLE_VOTER+` SET voteRecords = $1 where openid = $2`, voteRecords, openID)
	if err != nil {
		return
	}

	return tx.Commit()
}

func (d DB) ListVoters(key interface{}, index, size int) (*VoterList, error) {
	var sql string = `SELECT voterid, COALESCE(name, ''), COALESCE(image, ''), votedcount FROM ` + TABLE_VOTER + `
		WHERE followed AND registed AND imagecached  %s ORDER BY votedcount DESC LIMIT %d OFFSET %d`
	var countSQL string = `SELECT count(*) FROM ` + TABLE_VOTER + ` WHERE imagecached AND registed %s`
	var sqlCondition string
	if key == nil {
		sqlCondition = ""
	} else {
		switch value := key.(type) {
		case int:
			sqlCondition = fmt.Sprintf(`AND voterid = %d`, value)
		case string:
			sqlCondition = fmt.Sprintf(`AND name like '%%%s%%'`, value)
		}
	}

	rows, err := d.Query(fmt.Sprintf(sql, sqlCondition, size, index-1))
	if err != nil {
		return nil, err
	}

	vs := []Voter{}
	for rows.Next() {
		v := Voter{}
		if err := rows.Scan(&v.ID, &v.Name, &v.Image, &v.VotedCount); err != nil {
			glog.Errorf("list voters scan row err %v\n", err)
			continue
		}
		vs = append(vs, v)
	}

	var tatol int
	d.QueryRow(fmt.Sprintf(countSQL, sqlCondition)).Scan(&tatol)
	l := VoterList{
		Index:      index,
		Size:       size,
		TotalPages: tatol / 100,
		PageData:   vs,
	}

	if tatol%100 > 0 {
		l.TotalPages++
	}

	return &l, nil
}

func (d DB) GetVoter(openid string) (*Voter, error) {
	v := Voter{}

	if err := d.QueryRow(`SELECT voterid, COALESCE(name, ''), COALESCE(image, ''),
		COALESCE(company, ''), COALESCE(mobile, ''), COALESCE(votedcount, 0),　
		followed, registed, imageCached FROM `+TABLE_VOTER+` WHERE openid = $1`, openid).
		Scan(&v.ID, &v.Name, &v.Image, &v.Company, &v.Mobile, &v.VotedCount,
			&v.followed, &v.registed, &v.imageCached); err != nil {
		if err == sql.ErrNoRows {
			return nil, ERR_NOT_FOLLOW
		}
		return nil, err
	}

	if !v.registed {
		return nil, ERR_NOT_REGISTER
	}

	return &v, nil
}

func (d DB) Follow(e *event.Event)( err error) {
	fmt.Printf("--------------follow event-------------------------> %v\n", *e)
	_, err = d.Exec(`INSERT INTO ` + TABLE_VOTER+ ` (openid, followed) VALUES($1, TRUE)
		ON CONFLICT(openid) DO UPDATE SET openid=EXCLUDED.openid, followed=EXCLUDED.followed`, e.From)
	fmt.Println("--------------------------------------->", err)
	return
}

func (d DB) UnFollow(e *event.Event) (err error) {
	fmt.Printf("-------------------unfollow event--------------------> %v\n", *e)
	_, err = d.Exec(`UPDATE ` + TABLE_VOTER+ ` SET followed = FALSE WHERE openid = $1`, e.From)
	fmt.Println("--------------------------------------->", err)
	return
}

func (d DB) updateVoterImageStatus(image string) (err error) {
	_, err = d.Exec(`UPDATE `+TABLE_VOTER+` SET image = concat(image, '.jpg') , imageCached = TRUE where image = $1`, image)
	return
}

func hasVoteRight(records []int64) bool {
	if len(records) >= 1 {
		sort.Sort(Int64Slice(records))
		return time.Now().Unix()-records[2] >= 3600*24
	} else {
		return true
	}

}
