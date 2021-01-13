package cron

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	support, _ := Read()
	fmt.Sprintf("%v", support)
}
