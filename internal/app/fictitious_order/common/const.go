package common

/*
 * @Content : common
 * @Author  : LiJunDong
 * @Time    : 2022-09-13$
 */

const (
	CookieSessionSaltKey = "cookie_session_salt"
	UserPasswordSaltKey = "password_salt"
)

const (
	RegisterCodeMin = 1000
	RegisterCodeMax = 9999
	RegisterCodeExpire = 60
)

const (
	RegisterCodeOnOffConfig = "register_code_on_off"
	InvitationCodeOnOffConfig = "invitation_code_on_off"
)