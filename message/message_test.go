package message

import (
	"fmt"
	"testing"
)

func TestEnabled(t *testing.T) {
	e := Enabled()
	fmt.Println(e)
}
