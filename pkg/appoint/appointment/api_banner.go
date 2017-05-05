package appointment

import (
	"bjdaos/pegasus/pkg/appoint/db"
	"fmt"
)

//pos INTEGER,
//imageUrl VARCHAR(30),
//redirectUrl VARCHAR(30),
func GetBanners() ([]Banner, error) {
	sqlStr := fmt.Sprintf("SELECT pos,imageUrl,redirectUrl FROM %s", TABLE_BANNER)
	bs := make([]Banner, 0)
	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	b := Banner{}
	for rows.Next() {
		if err = rows.Scan(b.Pos, b.IfShow, b.RedirectUrl); err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return bs, nil
}
