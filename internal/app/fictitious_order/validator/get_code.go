package validator

import (
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
type GetCodeReq struct {
	To string `form:"to" binding:"required" label:"接收方"`
}

func CheckGetCodeParam(c *gin.Context) {
	this := rgrequest.Get(c)
	var param GetCodeReq
	err := c.ShouldBind(&param)
	if err != nil {
		errMsg, _ := rgrouter.Error(err)
		this.Response.ReturnError(-500, nil, errMsg)
		return
	}
	if err = checkTo(param); err != nil {
		this.Response.ReturnError(-500, nil, err.Error())
		return
	}
	this.Param = param
	c.Next()
}

func checkTo(param GetCodeReq) (err error) {
	if rgconfig.GetInt(common.RegisterGetCodeType) == common.SendTypePhone {
		return CheckPhone(param.To)
	} else {
		return CheckEmail(param.To)
	}
}

// HighFrequencyGetCodeLock <LiJunDong : 2022-11-06 16:03:57> --- TODO 未实现 控制请求频率，获取验证码频率限制
func HighFrequencyGetCodeLock(c *gin.Context) {
	this := rgrequest.Get(c)
	this.Log.Info("HighFrequencyRequestLock ---- Before")
	c.Next()
	this.Log.Info("HighFrequencyRequestLock ---- After")
}
