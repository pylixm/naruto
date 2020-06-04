package common

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
)

var (
	DefaultPageSize int64 = 20
)

func init() {
	DefaultPageSize = beego.AppConfig.DefaultInt64("page_size", 20)
}

func GetMd5(s string) string {
	b := []byte(s)
	md5Ctx := md5.New()
	md5Ctx.Write(b)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
