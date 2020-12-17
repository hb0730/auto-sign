package config

import (
	"auto-sign/util"
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {

	rMap := Read()
	fmt.Println(rMap)
}

func TestRedStruct(t *testing.T) {
	result, err := RedStruct()
	if err != nil {
		util.ErrorF("出错 %v \n", err)
	}
	fmt.Println(result)
}
