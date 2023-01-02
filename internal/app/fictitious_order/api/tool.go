package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/util/rgtime"
	"time"
	"video-admin/internal/app/fictitious_order/common"
	"video-admin/internal/app/fictitious_order/validator"
	"video-admin/pkg/mysql"
	"video-admin/pkg/mysql/video"
	"video-admin/pkg/random_code"
)

type RegisterCodeClient struct {
	This       *rgrequest.Client
	Code       string
	Msg        string
	To         string
	ExpireTime int64
}

func GetCodeHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	req := this.Param.(validator.GetCodeReq)
	client := RegisterCodeClient{
		This: this,
		To:   req.Phone,
	}
	if err := client.SetCode(); err != nil {
		common.ReturnErrorAndLog(this, -4000, "获取验证码失败", err)
		return
	}
	if err := client.Send(); err != nil {
		common.ReturnErrorAndLog(this, -4001, "获取验证码失败", err)
		return
	}
	if err := client.Save(); err != nil {
		common.ReturnErrorAndLog(this, -4002, "获取验证码失败", err)
		return
	}
	this.Log.Info("验证码", req.Phone, client.Code)
	this.Response.ReturnSuccess(nil)
	return
}

func (m *RegisterCodeClient) SetCode() (err error) {
	code := random_code.GetCodeStr(common.RegisterCodeMin, common.RegisterCodeMax)
	m.Code = code
	return err
}

func (m *RegisterCodeClient) Send() (err error) {
	m.Msg = "你的验证码是" + m.Code + "，请勿泄露。"
	m.ExpireTime = rgtime.NowTimeInt() + common.RegisterCodeExpire
	return err
}

func (m *RegisterCodeClient) Save() (err error) {
	expireTime := time.Unix(m.ExpireTime, 64)
	model := video.ValidCode{
		Code:    m.Code,
		Phone:   m.To,
		Status:  1,
		MsgType: 1,
		Msg:     m.Msg,
		ExpireTime: mysql.Time(expireTime),
	}
	return model.Create(m.This)
}
