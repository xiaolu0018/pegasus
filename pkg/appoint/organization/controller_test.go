package organization

import (
	"bjdaos/pegasus/pkg/appoint/db"
	"fmt"
	"github.com/lib/pq"
	"testing"
)

func TestAutoDeleteOverdueOffdays(t *testing.T) {

	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		t.Fatal(err)
	}

	basics, err := GetAllGetConfigBasics()
	if err != nil {
		t.Fatal(err)
	}
	for _, basic := range basics {
		sqlStr := fmt.Sprintf(`UPDATE %s SET offdays = $1 WHERE org_code = '%s'`, TABLE_ORG_CON_BASIC, basic.Org_Code)

		if _, err := db.GetDB().Exec(sqlStr, pq.Array(deleteOverdueOffdays(basic.OffDays))); err != nil {

			t.Fatal(err)
		}
	}
}
