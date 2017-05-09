package appointment

import (
	"bjdaos/pegasus/pkg/appoint/db"
	"fmt"
)

func GetBanners() ([]Banner, error) {
	sqlStr := fmt.Sprintf("SELECT pos,imageUrl,redirectUrl FROM %s", T_BANNER)
	bs := make([]Banner, 0)
	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	b := Banner{}
	for rows.Next() {
		if err = rows.Scan(&b.Pos, &b.ImageUrl, &b.RedirectUrl); err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return bs, nil
}
