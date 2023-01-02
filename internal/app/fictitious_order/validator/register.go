package validator

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/core/rgrouter"
	"video-admin/internal/app/fictitious_order/common"
	"video-admin/internal/app/fictitious_order/middleware/check_valid_code"
)

/*
 * @Content : validator
 * @Author  : LiJunDong
 * @Time    : 2022-09-14$
 */
type RegisterReq struct {
	Phone          string `form:"phone" binding:"required"`
	ValidCode      string `form:"valid_code" binding:"required"`
	ValidCodeID    int
	InvitationCode string `form:"invitation_code" binding:"required"`
	Password       string `form:"password" binding:"required"`
	SurePassword   string `form:"sure_password" binding:"required"`
}

func CheckRegisterParam(c *gin.Context) {
	this := rgrequest.Get(c)
	var param RegisterReq
	err := c.ShouldBind(&param)
	if err != nil {
		errMsg, _ := rgrouter.Error(err)
		common.ReturnErrorAndLog(this, -500, "参数错误", errors.New(errMsg))
		return
	}
	if !param.checkSurePassword() {
		common.ReturnErrorAndLog(this, -500, "确认密码不匹配", err)
		return
	}

	this.Param = param
	c.Next()
}

func (m RegisterReq) checkSurePassword() (res bool) {
	if m.Password != m.SurePassword {
		return false
	}
	return true
}

func ValidateRegisterCode(c *gin.Context) {
	if !rgconfig.GetBool(common.RegisterCodeOnOffConfig) {
		c.Next()
		return
	}
	this := rgrequest.Get(c)
	req := this.Param.(RegisterReq)
	id, res, err := check_valid_code.CheckValidRegisterCode(this, req.Phone, req.ValidCode)
	if err != nil {
		common.ReturnErrorAndLog(this, -500, "验证码验证失败", err)
		return
	}
	if !res {
		common.ReturnErrorAndLog(this, -500, "验证码错误", err)
		return
	}
	req.ValidCodeID = id
	this.Param = req
	c.Next()
}

// TODO <LiJunDong : 2022-11-06 17:30:02> --- 邀请码验证

func ValidateInvitationCode(c *gin.Context) {
	if !rgconfig.GetBool(common.InvitationCodeOnOffConfig) {
		c.Next()
		return
	}
	c.Next()
}
