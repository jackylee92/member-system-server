package common

/*
 * @Content : common
 * @Author  : LiJunDong
 * @Time    : 2022-09-13$
 */

const (
	UserPasswordSalt           = "ljd"                        // 用戶密码生产 盐
	UserTokenJWTSalt           = "ljd@jacky@0230@balabala*&!" // JWT token生成 盐
	UserTokenJWTExpireDuration = 600                          // JWT token过期时间 秒
	//  <LiJunDong : 2023-01-11 11:23:06> --- 不能和验证码状态混淆，jwt和validCode状态分开表示
	//  <LiJunDong : 2023-01-11 13:51:10> --- type是token的类型 status是状态
	JWTTokenTypeLogin            = 1   // token中type：登录
	JWTTokenTypeForget           = 2   // token中type：忘记密码
	JWTTokenTypeRegister         = 3   // token中type：注册
	JWTTokenForgetCodeNoUse      = 1   // 忘记密码的token未被核销
	JWTTokenForgetCodeUsed       = 2   // 忘记密码的token已被核销
	JWTTokenRegisterCodeNoUse    = 1   // 注册的token未被核销
	JWTTokenRegisterCodeUsed     = 2   // 注册的token已被核销
	JWTTokenForgetExpireDuration = 600 // 忘记密码的token过期时间
)

const (
	RegisterCodeMin    = 1000 // 注册验证码范围
	RegisterCodeMax    = 9999 // 注册验证码范围
	RegisterCodeExpire = 600  // 注册验证码过期时间
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

// <LiJunDong : 2023-01-11 16:21:12> --- 生成随机码需要的集合
var AlphanumericSet = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

// <LiJunDong : 2023-01-11 16:21:25> --- 推荐码长度
const InvitationCodeLength = 8
