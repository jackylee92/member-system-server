package validator

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/core/rgrouter"
	"member-system-server/internal/app/fictitious_order/common"
)

/*
 * @Content : validator
 * @Author  : LiJunDong
 * @Time    : 2022-09-14$
 */
type RegisterReq struct {
	Phone          string `form:"phone" binding:"required" label:"手机号"`
	ValidCode      string `form:"valid_code" label:"验证码"`
	ValidCodeID    int
	InvitationCode string `form:"invitation_code" binding:"required" label:"邀请码"`
	Password       string `form:"password" binding:"required" label:"密码"`
	SurePassword   string `form:"sure_password" binding:"required" label:"确认密码"`
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
	if !param.checkSurePassword() {
		common.ReturnErrorAndLog(this, -500, "确认密码不匹配", err)
		return
	}
	validCodeId, err := validateRegisterCode(this, param)
	if err != nil {
		common.ReturnErrorAndLog(this, -500, err.Error(), err)
		return
	}
	param.ValidCodeID = validCodeId
	this.Param = param
	c.Next()
}

func (m RegisterReq) checkSurePassword() (res bool) {
	if m.Password != m.SurePassword {
		return false
	}
	return true
}

// validateRegisterCode <LiJunDong : 2023-01-06 16:35:15> --- 验证码
func validateRegisterCode(this *rgrequest.Client, param RegisterReq) (validCodeId int, err error) {
	if !rgconfig.GetBool(common.RegisterCodeOnOffConfig) {
		return validCodeId, err
	}
	return CheckValidCode(this, common.SendTypePhone, param.Phone, param.ValidCode)
}

// TODO <LiJunDong : 2022-11-06 17:30:02> --- 邀请码验证

func ValidateInvitationCode(c *gin.Context) {
	if !rgconfig.GetBool(common.InvitationCodeOnOffConfig) {
		c.Next()
		return
	}
	c.Next()
}
