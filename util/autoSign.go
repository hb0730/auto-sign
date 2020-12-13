package util

import (
	"crypto"
	"encoding/hex"
	"github.com/go-rod/rod/lib/proto"
)

func GetMd5(params string) string {
	md5Ctx := crypto.MD5.New()
	md5Ctx.Write([]byte(params))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func ConvertCookies(cookies Cookies, domain string) []*proto.NetworkCookie {
	c := make([]*proto.NetworkCookie, 0)
	for k, v := range cookies {
		c = append(c, &proto.NetworkCookie{
			Name:   k,
			Value:  v,
			Domain: domain,
		})
	}
	return c
}
