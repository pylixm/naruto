package filters

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"naruto/common"
	"naruto/models"
	"strings"
)

func Err401(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(401)
	ctx.Output.Body([]byte("{\"code\": \"401\", \"message\": \"未登录，或登录已失效\", \"data\": \"token expired\"}"))
}

func AuthUserToken(ctx *context.Context) {
	jt := common.JWTToken{}
	authtoken := strings.TrimSpace(ctx.Request.Header.Get("Authorization"))
	valido, _ := jt.ValidateToken(authtoken)
	if !valido && ctx.Request.RequestURI != "/login" {
		Err401(ctx)
	}
}

func AuthUser(ctx *context.Context) {
	u := ctx.GetSession("UserInfo")
	uInfo := models.NarutoUser{}
	json.Unmarshal([]byte(fmt.Sprintf("%v", u)), &uInfo)

}
