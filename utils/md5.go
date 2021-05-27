package utils

import (
	"crypto"
	"encoding/hex"
)

//GetMd5 通过Md5摘要
func GetMd5(params string) string {
	md5Ctx := crypto.MD5.New()
	md5Ctx.Write([]byte(params))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
