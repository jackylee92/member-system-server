package common

/*
 * @Content : common
 * @Author  : LiJunDong
 * @Time    : 2022-09-13$
 */

const (
	UserPasswordSaltKey = "password_salt"
)

const (
	RegisterCodeMin    = 1000
	RegisterCodeMax    = 9999
	RegisterCodeExpire = 60
)

const (
	RegisterCodeOnOffConfig   = "register_code_on_off"   // 手机验证码
	InvitationCodeOnOffConfig = "invitation_code_on_off" // 注册推荐码
)

const (
	SendTypePhone = 2
	SendTypeEmail = 1
)
