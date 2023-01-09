package validator

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/core/rgrouter"
)

/*
 * @Content : validator
 * @Author  : LiJunDong
 * @Time    : 2022-09-14$
 */
type ForgetNewPasswordReq struct {
	To           string `form:"to" binding:"required" label:"手机号/邮箱"`
	UserId       int
	NewPassword  string `form:"new_password" binding:"required" label:"新密码"`
	SurePassword string `form:"sure_password" binding:"required" label:"确认密码"`
}

func CheckForgetNewPasswordParam(c *gin.Context) {
	this := rgrequest.Get(c)
	var param ForgetNewPasswordReq
	err := c.ShouldBind(&param)
	if err != nil {
		errMsg, _ := rgrouter.Error(err)
		this.Response.ReturnError(-500, nil, errMsg)
		return
	}

	if err = param.checkSurePassword(); err != nil {
		this.Response.ReturnError(-500, nil, err.Error())
		return
	}
	this.Param = param
	c.Next()
}

func (m *ForgetNewPasswordReq) checkSurePassword() (err error) {
	if m.NewPassword != m.SurePassword {
		return errors.New("确认密码错误")
	}
	return err
}
