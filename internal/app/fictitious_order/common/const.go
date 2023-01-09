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
	RegisterGetCodeType       = "register_get_code_type" // 注册获取验证码方式 1, 2
)

const (
	SendTypePhone = 2
	SendTypeEmail = 1
)

const (
	DefaultUserAvatarUrl    = "https://avatars.githubusercontent.com/u/24886757?v=4"
	DefaultUserIntroduction = "这家伙很拽，啥都没说！"
)

const (
	HighFrequencyGetCodeTime     = 2   // 距离上次请求间隔HighFrequencyGetCodeTime秒为一次高频请求
	HighFrequencyGetCodeLockTime = 300 // 连续HighFrequencyGetCodeTimes次高频请求后需要等待HighFrequencyGetCodeLockTime秒后可再次请求
	HighFrequencyGetCodeTimes    = 3   // 连续HighFrequencyGetCodeTimes次高频请求后需要等待HighFrequencyGetCodeLockTime秒后可再次请求
)
