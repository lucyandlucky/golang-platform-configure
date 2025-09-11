package service

import (
	"configure/api/configure/errors"
	"configure/internal/conf"
	"encoding/base64"
	"encoding/json"
	"github.com/forgoer/openssl"
	"github.com/limes-cloud/kratosx"
)

type User struct {
	conf *conf.Config
}

func NewUser(conf *conf.Config) *User {
	return &User{
		conf: conf,
	}
}

type Password struct {
	Password string `json:"password"`
	Time     int64  `json:"time"`
}

func (us *User) Login(ctx kratosx.Context, username, password string) (string, error) {
	pwByte, err := base64.StdEncoding.DecodeString(password)

	if err != nil {
		return "", errors.ParamsError()
	}

	decryptData, err := openssl.RSADecrypt(pwByte, ctx.Loader("login"))

	if err != nil {
		return "", errors.ParamsError()
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return "", errors.PasswordError()
	}

	//if time.Now().UnixMilli()-pw.Time > 10*1000 {
	//	return "", errors.PasswordExpireError()
	//}

	if us.conf.Author.AdminPassword != pw.Password || us.conf.Author.AdminUser != username {
		return "", errors.PasswordError()
	}

	token, err := ctx.JWT().NewToken(map[string]any{"userId": 1, "roleKeyword": "superAdmin"})

	if err != nil {
		return "", errors.SystemError(err.Error())
	}

	return token, nil
}
