package message

import "github.com/jackylee92/rgo/core/rgglobal/rgmessage"

const (
	TEST        = 201
	AppLockCode = -1000
	NoLoginCode = -2000
	NoApi       = -4040
	NoUserId    = -3000
	UserInfoErr = -3001
)

const (
	testMsg     = "测试返回"
	appLockMsg  = "系统已关闭"
	noLoginMsg  = "您还未登录"
	noApi       = "地址错误"
	noUserId    = "用户ID不存在"
	userInfoErr = "用户信息异常"
)

var data = map[int64]string{
	TEST:        testMsg,
	AppLockCode: appLockMsg,
	NoLoginCode: noLoginMsg,
	NoApi:       noApi,
	NoUserId:    noUserId,
}

/*
* @Content : 注入语言包
* @Param   :
* @Return  :
* @Author  : LiJunDong
* @Time    : 2022-03-10
 */
func init() {
	rgmessage.InitAppMsg(data)
}
