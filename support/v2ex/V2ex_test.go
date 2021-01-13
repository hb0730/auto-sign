package v2ex

import (
	"fmt"
	"testing"
)

func TestV2ex_DoSupport(t *testing.T) {

}

func TestV2ex_Read(t *testing.T) {
	v := V2ex{}
	job, _ := v.Read()
	fmt.Sprintf("%v", job)
}
