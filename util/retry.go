package util

import (
	"github.com/go-rod/rod"
	"time"
)

//尝试机制
func Retry(page *rod.Page, autoSign AutoSign, num int) {
	err := rod.Try(func() {
		autoSign.Checking(page)
	})
	if err != nil && num > 0 {
		for {
			num--
			time.Sleep(time.Duration(3) * time.Second)
			Retry(page, autoSign, num)
		}
	} else if err != nil {
		panic(err)
	}
}
