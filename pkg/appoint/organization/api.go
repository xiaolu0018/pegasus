package organization

import (
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/db"
	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
	"fmt"
	"github.com/lib/pq"
)

func ListDBOrgs() ([]types.Organization, error) {
	rows, err := db.GetDB().Query("SELECT org_code, id, name FROM " + TABLE_ORG)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := []types.Organization{}
	var code, id, name string
	for rows.Next() {
		if err = rows.Scan(&code, &id, &name); err != nil {
			return nil, err
		}

		l = append(l, types.Organization{ID: id, Code: code, Name: name})

	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return l, nil
}

func CreateOrg(org types.Organization) error {
	_, err := db.GetDB().Exec("INSERT INTO "+TABLE_ORG+"(ORG_CODE, ID, NAME,PHONE,IMAGEURL, DETAILSURL) VALUES ($1, $2, $3,$4,$5)",
		org.Code, org.ID, org.Name, org.Phone, org.ImageUrl, org.DetailsUrl)
	return err
}

func UpdateOrg(org types.Organization) error {
	_, err := db.GetDB().Exec("UPDATE "+TABLE_ORG+" SET ID = $1, NAME = $2 ,PHONE = $3 WHERE ORG_CODE = $4",
		org.ID, org.Name, org.Phone, org.Code)
	return err
}

func DeleteOrg(org types.Organization) error {
	_, err := db.GetDB().Exec("UPDATE "+TABLE_ORG+" SET DELETED = true WHERE ORG_CODE = $1", org.Code)
	return err
}

func CreateOrgs(orgs []types.Organization) error {
	for _, org := range orgs {
		if err := CreateOrg(org); err != nil {
			return err
		}
	}
	return nil
}

func UpdateOrgs(orgs []types.Organization) error {
	for _, org := range orgs {
		if err := UpdateOrg(org); err != nil {
			return err
		}
	}
	return nil
}

func DeleteOrgs(orgs []types.Organization) error {
	for _, org := range orgs {
		if err := DeleteOrg(org); err != nil {
			return err
		}
	}
	return nil
}

func codesArray(orgs []types.Organization) []string {
	l := []string{}
	for _, org := range orgs {
		l = append(l, org.Code)
	}
	return l
}

func ListOrganizations(page_index, page_size int) ([]Organization, error) {
	sql := `SELECT a.id, a.ORG_CODE, a.name, COALESCE(b.capacity, 0) as capacity, COALESCE(b.warnnum, 0) as warnnum, b.offdays, b.avoidnumbers FROM %s a LEFT JOIN %s b ON a.org_code = b.org_code WHERE NOT deleted ORDER BY id LIMIT %d OFFSET %d`
	sql = fmt.Sprintf(sql, TABLE_ORG, TABLE_ORG_CON_BASIC, page_size, page_index)
	rows, err := db.GetDB().Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := []Organization{}
	org := Organization{}
	bc := Config_Basic{}
	offDays := pq.StringArray{}
	avoidNumbers := pq.Int64Array{}
	for rows.Next() {
		if err = rows.Scan(&org.ID, &org.Code, &org.Name,
			&bc.Capacity, &bc.WarnNum, &offDays, &avoidNumbers); err != nil {
			return nil, err
		}

		bc.OffDays = ([]string)(offDays)
		bc.AvoidNumbers = ([]int64)(avoidNumbers)
		org.BasicCon = bc

		l = append(l, org)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return l, nil
}

func ListOrganizationsForWC() ([]Organization, error) {
	sql := `SELECT org_code, name, phone, imageurl,detailsUrl FROM %s`
	sql = fmt.Sprintf(sql, TABLE_ORG)
	rows, err := db.GetDB().Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := []Organization{}
	org := Organization{}
	for rows.Next() {
		if err = rows.Scan(&org.Code, &org.Name, &org.Phone, &org.ImageUrl, &org.DetailsUrl); err != nil {
			return nil, err
		}

		l = append(l, org)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return l, nil
}
func GetOrgByCode(code string) (*Organization, error) {
	org := Organization{}
	if err := db.GetDB().QueryRow(`SELECT org_code, id, name FROM `+TABLE_ORG+` WHERE org_code = $1`, code).
		Scan(&org.Code, &org.ID, &org.Name); err != nil {

		return nil, err
	}

	return &org, nil
}
