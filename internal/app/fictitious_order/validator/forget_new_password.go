package validator

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/core/rgrouter"
	"member-system-server/internal/app/fictitious_order/api/user"
	"member-system-server/internal/app/fictitious_order/common"
)

/*
 * @Content : validator
 * @Author  : LiJunDong
 * @Time    : 2022-09-14$
 */
type ForgetNewPasswordReq struct {
	To           string `json:"to" binding:"required" label:"手机号/邮箱"`
	UserId       int
	NewPassword  string `json:"new_password" binding:"required" label:"新密码"`
	SurePassword string `json:"sure_password" binding:"required" label:"确认密码"`
}

// TODO <LiJunDong : 2023/1/10 0:08> --- 需要验证 找回密码-发送验证码成功返回的token
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
	userInfo := user.CheckForgetAuthorization(this, common.JWTTokenForgetCodeUsed)
	if userInfo.UserId == 0 {
		common.ReturnErrorAndLog(this, -500, "验证码错误，请重新获取", nil)
		return
	}
	param.UserId = userInfo.UserId
	this.Param = param
	c.Next()
}

func (m *ForgetNewPasswordReq) checkSurePassword() (err error) {
	if m.NewPassword != m.SurePassword {
		return errors.New("确认密码错误")
	}
	return err
}
