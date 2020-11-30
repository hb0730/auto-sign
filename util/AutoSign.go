package util

import (
	"crypto"
	"encoding/hex"
)

func GetMd5(params string) string {
	md5Ctx := crypto.MD5.New()
	md5Ctx.Write([]byte(params))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
