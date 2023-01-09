package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/core/rgrouter"
	"member-system-server/internal/app/fictitious_order/common"
)

/*
 * @Content : validator
 * @Author  : LiJunDong
 * @Time    : 2022-09-14$
 */
type ForgetCheckCodeReq struct {
	To          string `json:"to" binding:"required" label:"手机号/邮箱"`
	ValidCode   string `json:"code" binding:"required" label:"验证码"`
	ValidCodeID int
	SendType    int8
}

// TODO <LiJunDong : 2023/1/10 0:07> --- 需要验证找回密码-获取验证码成功后返回的token
func CheckForgetCheckCodeParam(c *gin.Context) {
	this := rgrequest.Get(c)
	var param ForgetCheckCodeReq
	err := c.ShouldBind(&param)
	if err != nil {
		errMsg, _ := rgrouter.Error(err)
		this.Response.ReturnError(-500, nil, errMsg)
		return
	}

	sendType, err := getCodeAcceptType(param.To)
	if err != nil {
		this.Response.ReturnError(-500, nil, err.Error())
		return
	}
	param.SendType = sendType

	validCodeId, err := param.validateForgetCheckCode(this)
	if err != nil {
		common.ReturnErrorAndLog(this, -500, err.Error(), err)
		return
	}
	param.ValidCodeID = validCodeId
	this.Param = param
	c.Next()
}

// validateForgetCheckCode <LiJunDong : 2023-01-06 16:35:15> --- 验证码
func (m *ForgetCheckCodeReq) validateForgetCheckCode(this *rgrequest.Client) (validCodeId int, err error) {
	return checkValidCode(this, m.SendType, m.To, m.ValidCode)
}
