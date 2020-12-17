package main

import (
	"auto-sign/util"
	config "auto-sign/yml"
	"testing"
)

func TestMain_test(T *testing.T) {
	main()
}
func TestMain_readFile(T *testing.T) {
	autoSign, err := config.RedStruct()
	if err != nil {
		util.ErrorF("%v \n", err)
	}
	for k, v := range autoSign.Cron {
		util.InfoF("key: %v ,value: %v \n", k, v)
	}
}
