package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"member-system-server/internal/app/fictitious_order/common"
	"strconv"
	"time"
)

/*
 * @Content : jwt
 * @Author  : LiJunDong
 * @Time    : 2022-11-29$
 */

type LoginData struct {
	Login     bool
	UserId    string // 用string，避免int莫名其妙变成float类型
	Username  string
	Typ       string // 1：登录 2：忘记密码
	Status    string
	ValidCode string
	ValidId   string
	jwt.StandardClaims
}

// 生成 jwt token
func GetToken(this *rgrequest.Client, claims LoginData) (string, error) {
	this.Log.Debug("GetToken", claims)
	claimsType, _ := strconv.Atoi(claims.Typ)
	if claimsType == 0 {
		return "", errors.New("类型错误")
	}
	if claimsType == common.JWTTokenTypeLogin {
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(common.UserTokenJWTExpireDuration) * time.Second).Unix() // 过期时间
	}
	if claimsType == common.JWTTokenTypeForget {
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(common.JWTTokenForgetExpireDuration) * time.Second).Unix() // 过期时间
	}
	claims.StandardClaims.Issuer = rgconfig.GetStr("sys_app_name")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(common.UserTokenJWTSalt))
	if err != nil {
		return "", fmt.Errorf("生成token失败:%v", err)
	}
	return signedToken, nil
}

func secret() jwt.Keyfunc { //按照这样的规则解析
	return func(t *jwt.Token) (interface{}, error) {
		return []byte(common.UserTokenJWTSalt), nil
	}
}

// 解析token
func ParseToken(this *rgrequest.Client, token string) (loginData *LoginData, err error) {
	loginData = &LoginData{}
	tokenTmp, err := jwt.Parse(token, secret())
	if err != nil {
		this.Log.Error("jwt.Parse", err, token)
		return loginData, err
	}
	claim, ok := tokenTmp.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("解析错误")
		return loginData, err
	}
	if !tokenTmp.Valid {
		err = errors.New("Token无效")
		return loginData, err
	}
	this.Log.Info("Token解析", claim)
	loginData.Username = claim["Username"].(string)
	loginData.UserId = claim["UserId"].(string)
	loginData.Login = claim["Login"].(bool)
	loginData.Typ = claim["Typ"].(string)
	loginData.Status = claim["Status"].(string)
	loginData.ValidCode = claim["ValidCode"].(string)
	loginData.ValidId = claim["ValidId"].(string)
	return loginData, err
}
