package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"member-system-server/internal/app/fictitious_order/api/valid_code"
	"member-system-server/internal/app/fictitious_order/common"
	"member-system-server/internal/app/fictitious_order/validator"
)

func GetCodeHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	req := this.Param.(validator.GetCodeReq)
	client := valid_code.ValidCodeClient{
		This: this,
		To:   req.Phone,
		Typ:  common.SendTypePhone,
	}
	if err := client.GetCode(); err != nil {
		common.ReturnErrorAndLog(this, -4000, "获取验证码失败", err)
		return
	}
	this.Response.ReturnSuccess(nil)
	return
}
