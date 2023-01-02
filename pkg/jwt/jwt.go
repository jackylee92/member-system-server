package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"time"
)

const (
	jwtExpireDurationConfig = "jwt_expire_duration"
	jwtSaltConfig           = "jwt_salt"
)

/*
 * @Content : jwt
 * @Author  : LiJunDong
 * @Time    : 2022-11-29$
 */

type LoginData struct {
	Login    bool
	UserId   int
	Username string
	jwt.StandardClaims
}

//生成 jwt token
func GetToken(this *rgrequest.Client, claims LoginData) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(rgconfig.GetInt(jwtExpireDurationConfig)) * time.Second).Unix() // 过期时间
	claims.StandardClaims.Issuer = "ljd"                                                                                           // 签发人
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(rgconfig.GetStr(jwtSaltConfig)))
	if err != nil {
		return "", fmt.Errorf("生成token失败:%v", err)
	}
	return signedToken, nil
}

//验证jwt token
func ParseToken(this *rgrequest.Client, tokenStr string) (*LoginData, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &LoginData{}, func(token *jwt.Token) (i interface{}, err error) { // 解析token
		return rgconfig.GetStr(jwtSaltConfig), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*LoginData); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
