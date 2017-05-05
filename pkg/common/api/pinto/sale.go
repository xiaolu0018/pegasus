package pinto

import (
	"bjdaos/pegasus/pkg/common/types"
	"database/sql"
	"fmt"
	"strings"
)

//Sale_Code string
//Sale_OrgPrice float64
//Sale_SellPrice float64
//Sale_Discount float64
//LimitSex int
//Sale_Name string
//Brief_Name string
//Order_Position int
//IsValid int
//MnemSymbol string

func ListSales(db *sql.DB) ([]types.Sale, error) {
	rows, err := db.Query(`SELECT sale_code, sale_orgprice, sale_sellprice, sale_discount,
	  limit_sex, sale_name, brief_name, order_position, mnemsymbol, note
	  FROM sale where is_valid = 1 ORDER BY order_position`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := []types.Sale{}
	var s types.Sale
	for rows.Next() {
		s = types.Sale{}
		if err = rows.Scan(&s.Sale_Code, s.Sale_OrgPrice, s.Sale_SellPrice, s.Sale_Discount,
			s.LimitSex, s.Sale_Name, s.Brief_Name, s.Order_Position, s.MnemSymbol, s.Note); err == nil {
			l = append(l, s)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return l, nil
}

func GetSalesBySaleCodes(db *sql.DB, codes []string) ([]types.Sale, error) {

	sqlCodes := make([]string, len(codes))
	for k, salecode := range codes {
		sqlCodes[k] = fmt.Sprintf(`'%s'`, salecode)
	}
	sqlStr := fmt.Sprintf("SELECT sale_code,sale_sellprice,sale_discount FROM sale WHERE sale_code IN (%s)", strings.Join(sqlCodes, ","))

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sale := types.Sale{}
	sales := []types.Sale{}

	for rows.Next() {
		if err = rows.Scan(&sale.Sale_Code, &sale.Sale_SellPrice, &sale.Sale_Discount); err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return sales, nil
}
