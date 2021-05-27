package utils

import (
	"fmt"
	"testing"
)

func TestConvertCookies(t *testing.T) {
	c := Cookies{"测试": "cc"}
	cook := ConvertCookies(c)
	fmt.Println(cook)
}
