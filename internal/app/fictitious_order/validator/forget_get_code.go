package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/core/rgrouter"
	"member-system-server/internal/app/fictitious_order/api/valid_code"
)

/*
 * @Content : validator
 * @Author  : LiJunDong
 * @Time    : 2022-09-14$
 */
type ForgetGetCodeReq struct {
	To       string `form:"to" binding:"required" label:"手机号/邮箱"`
	SendType int8
}

func CheckForgetGetCodeParam(c *gin.Context) {
	this := rgrequest.Get(c)
	var param ForgetGetCodeReq
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
	if err := param.HighFrequencyForgetGetCodeLock(this); err != nil {
		this.Response.ReturnError(-500, nil, err.Error())
		return
	}
	this.Param = param
	c.Next()
}

func (m *ForgetGetCodeReq) HighFrequencyForgetGetCodeLock(this *rgrequest.Client) (err error) {
	client := valid_code.ValidCodeClient{
		This: this,
		To:   m.To,
	}
	err = client.CheckCount()
	return err
}
