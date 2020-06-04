package common

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/dgrijalva/jwt-go"
)

type JWTToken struct {
	Username string
	Uid      int64
	Expires  int64
}

var (
	verifyKey  string
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
	ErrExpired = "token expired" // 令牌过期
	ErrOther   = "other error"   // 其他错误
)

func init() {
	// 配置文件配置生成jwt token 的「salt」token
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	verifyKey = appConf.String("jwt::salt")
}

func (jt JWTToken) GetToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: jt.Expires, // time.Unix(c.ExpiresAt, 0)
		Issuer:    jt.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(verifyKey))
	if err != nil {
		beego.Error("jwt genarate failed: ", err)
	}
	return tokenString, err
}

func (jt JWTToken) ValidateToken(tokenString string) (bool, error) {
	if tokenString == "" {
		return false, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})
	if token == nil {
		return false, errors.New(ErrInvalid)
	}
	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New(ErrInvalid)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New(ErrExpired)
		} else {
			return false, errors.New(ErrOther)
		}
	} else {
		return false, errors.New(ErrOther)
	}
}
