package appointment

import (
	"fmt"
	"testing"
)

func TestGetOffDaye(t *testing.T) {
	dbinit()
	result, err := GetOffDay("0001001")
	fmt.Println("result,err", result, err)
}
