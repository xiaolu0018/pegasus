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
)

var TABLE_VOTER string = "GO_WEICHAT_ACTIVITY_VOTER"
var SQL_TABLE = `
CREATE TABLE IF NOT EXISTS  GO_WEICHAT_ACTIVITY_VOTER (
   openid varchar(50) PRIMARY KEY,
   voterid SERIAL,
   registed BOOLEAN DEFAULT FALSE,
   name varchar(15),
   image varchar(300),
   company varchar(100),
   mobile varchar(20),
   declaration varchar(50),
   votedcount integer DEFAULT 0,
   voteRecords integer[],
   imageCached	BOOLEAN DEFAULT FALSE
);

CREATE OR REPLACE FUNCTION wc_Regist_Voter(p_openid varchar, p_name VARCHAR, p_image VARCHAR, p_company VARCHAR, p_mobile VARCHAR, p_declaration varchar, p_imageCached boolean) RETURNS VARCHAR AS $$
  DECLARE
    registedCount INTEGER;
    unregistedCount INTEGER;
  BEGIN
    SELECT count(*) from go_weichat_activity_voter where openid = p_openid AND registed into registedCount;
    IF  registedCount = 1 THEN
     RETURN 'exist';
    ELSE
        SELECT count(*) from go_weichat_activity_voter where openid = p_openid into unregistedCount;
        IF unregistedCount = 1 THEN
          UPDATE go_weichat_activity_voter SET name=p_name, image=p_image, company=p_company, mobile=p_mobile, declaration=p_declaration, registed = TRUE, imageCached = p_imageCached WHERE openid = p_openid;
        ELSE
          INSERT INTO go_weichat_activity_voter(openid, name, image, company, mobile, declaration, registed, imageCached) VALUES (p_openid, p_name, p_image, p_company, p_mobile, p_declaration, TRUE, p_imageCached);
        END IF;
      RETURN 'ok';
    END IF;
  END;
$$ LANGUAGE plpgsql;`

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
		if _, err:= d.Exec(`INSERT INTO ` +TABLE_VOTER+ `(openid) VALUES($1) ON CONFLICT(openid)
		DO UPDATE SET openid = EXCLUDED.openid `, users[id]); err != nil {
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
	imageCached := "false"
	if v.imageCached {
		imageCached = "true"
	}
	var result string
	var sql string = fmt.Sprintf(`SELECT wc_Regist_Voter('%s', '%s', '%s', '%s', '%s', '%s', %s)`,
		v.OpenID, v.Name, v.Image, v.Company, v.Mobile, v.Declaration, imageCached)
	if err := d.QueryRow(sql).Scan(&result); err != nil {
		fmt.Println(err)
	}
	switch result {
	case "exist":
		return errors.New("user exists")

	case "ok":
		return nil
	default:
		return fmt.Errorf("unknown err %s\n", result)
	}
}

func (d DB) Vote(openID, votedID string) (err error) {
	voteRecords := pq.Int64Array{}
	if err = d.QueryRow(`SELECT voteRecords FROM `+TABLE_VOTER+` WHERE openid = $1`, openID).Scan(&voteRecords); err != nil {
		return
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
	var sql string = `SELECT voterid, COALESCE(name, ''), COALESCE(image, ''), votedcount FROM ` + TABLE_VOTER + ` WHERE imagecached AND registed %s ORDER BY votedcount DESC LIMIT %d OFFSET %d`
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
	if err := d.QueryRow(`SELECT voterid, name, image, company, mobile, votedcount FROM `+TABLE_VOTER+` WHERE registed AND openid = $1`, openid).
		Scan(&v.ID, &v.Name, &v.Image, &v.Company, &v.Mobile, &v.VotedCount); err != nil {
		return nil, err
	}
	return &v, nil
}

func (d DB) updateVoterImageStatus(image string) (err error) {
	_, err = d.Exec(`UPDATE `+TABLE_VOTER+` SET image = concat(image, '.jpg') , imageCached = TRUE where image = $1`, image)
	return
}

func hasVoteRight(records []int64) bool {
	if len(records) >= 3 {
		sort.Sort(Int64Slice(records))
		return time.Now().Unix()-records[2] >= 3600*24
	} else {
		return true
	}

}
