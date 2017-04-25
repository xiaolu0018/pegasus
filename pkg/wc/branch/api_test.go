package branch

import (
	"testing"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
)

func TestBranch_Create(t *testing.T) {
	branch := Branch{}
	branch.Name= "北京第二体检中心"
	branch.Capacity = 10

	branch.Create(db.Branch())
}
