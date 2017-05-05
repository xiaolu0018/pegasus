package vote

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"github.com/lib/pq"
	"sort"
	"time"
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
   imageCached	boolean DEFAULT false
);
`

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

func (d DB) Init() (err error) {
	_, err = d.Exec(SQL_TABLE)
	return
}

func (d DB) Register(v *Voter) error {
	_, err := d.Exec(`INSERT INTO `+TABLE_VOTER+` (openid, name, image, company, mobile, declaration)
	VALUES($1,$2,$3,$4,$5,$6)`, v.OpenID, v.Name, v.Image, v.Company, v.Mobile, v.Declaration)
	return err
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
	var sql string = `SELECT voterid, name, image, votedcount FROM ` + TABLE_VOTER + ` WHERE imagecached %s ORDER BY votedcount DESC LIMIT %d OFFSET %d`
	var countSQL string = `SELECT count(*) FROM ` + TABLE_VOTER + ` WHERE imagecached %s`
	var sqlCondition string
	if key == nil {
		sqlCondition = ""
	} else {
		switch value := key.(type) {
		case int:
			sqlCondition = fmt.Sprintf(`AND voterid = %d`, value)
		case string:
			sqlCondition = fmt.Sprintf("AND name like '%s'", value)
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
		TotalPages: tatol / 8,
		PageData:   vs,
	}

	if tatol%8 > 0 {
		l.TotalPages++
	}

	return &l, nil
}

func (d DB) GetVoter(openid string) (*Voter, error) {
	v := Voter{}
	if err := d.QueryRow(`SELECT voterid, name, image, company, mobile, votedcount FROM `+TABLE_VOTER+` WHERE openid = $1`, openid).
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
	if len(records) >= 100 {
		sort.Sort(Int64Slice(records))
		return time.Now().Unix()-records[2] >= 3600*24
	} else {
		return true
	}

}
