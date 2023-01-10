package validator

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgconfig"
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
type RegisterReq struct {
	To               string `form:"to" binding:"required" label:"手机号/邮箱"`
	ValidCode        string `form:"valid_code" label:"验证码"`
	ValidCodeID      int
	InvitationCode   string `form:"invitation_code" binding:"required" label:"邀请码"`
	Password         string `form:"password" binding:"required" label:"密码"`
	SurePassword     string `form:"sure_password" binding:"required" label:"确认密码"`
	InvitationUserId int
}

func CheckRegisterParam(c *gin.Context) {
	this := rgrequest.Get(c)
	var param RegisterReq
	err := c.ShouldBind(&param)
	if err != nil {
		errMsg, _ := rgrouter.Error(err)
		common.ReturnErrorAndLog(this, -500, errMsg, errors.New(errMsg))
		return
	}
	if !param.checkRegisterSurePassword() {
		common.ReturnErrorAndLog(this, -500, "确认密码不匹配", err)
		return
	}
	err = param.validateRegisterCode(this)
	if err != nil {
		common.ReturnErrorAndLog(this, -500, err.Error(), err)
		return
	}
	if err = param.ValidateRegisterInvitationCode(this); err != nil {
		common.ReturnErrorAndLog(this, -500, err.Error(), err)
		return
	}
	this.Log.Info("register", param)
	this.Param = param
	c.Next()
}

func (m *RegisterReq) checkRegisterSurePassword() (res bool) {
	if m.Password != m.SurePassword {
		return false
	}
	return true
}

// validateRegisterCode <LiJunDong : 2023-01-06 16:35:15> --- 验证码
func (m *RegisterReq) validateRegisterCode(this *rgrequest.Client) (err error) {
	if !rgconfig.GetBool(common.RegisterCodeOnOffConfig) {
		return err
	}
	validCodeId, _, err := checkValidCode(this, int8(rgconfig.GetInt(common.RegisterGetCodeType)), m.To, m.ValidCode)
	if err != nil {
		return err
	}
	m.ValidCodeID = validCodeId
	return err
}

func (m *RegisterReq) ValidateRegisterInvitationCode(this *rgrequest.Client) (err error) {
	if !rgconfig.GetBool(common.InvitationCodeOnOffConfig) {
		return err
	}
	userInfo := user.Info{
		InvitationCode: m.InvitationCode,
	}
	err = userInfo.CheckInvitationCode(this)
	if err != nil {
		return err
	}
	m.InvitationUserId = userInfo.InvitationUserId
	return err
}
