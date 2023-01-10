package common

/*
 * @Content : common
 * @Author  : LiJunDong
 * @Time    : 2022-09-13$
 */

const (
	UserPasswordSalt             = "ljd"                        // 用戶密码生产 盐
	UserTokenJWTSalt             = "ljd@jacky@0230@balabala*&!" // JWT token生成 盐
	UserTokenJWTExpireDuration   = 600                          // JWT token过期时间 秒
	JWTTokenTypeLogin            = 1                            // token中type：注册
	JWTTokenTypeForget           = 2                            // token中type：忘记密码
	JWTTokenForgetCodeNoUse      = 1                            // 忘记密码的token未被核销
	JWTTokenForgetCodeUsed       = 2                            // 忘记密码的token已被核销
	JWTTokenForgetExpireDuration = 600                          // 忘记密码的token过期时间
)

const (
	RegisterCodeMin    = 1000 // 注册验证码范围
	RegisterCodeMax    = 9999 // 注册验证码范围
	RegisterCodeExpire = 600  // 注册验证码国企时间
)

const (
	RegisterCodeOnOffConfig   = "register_code_on_off"   // 注册验证码
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
