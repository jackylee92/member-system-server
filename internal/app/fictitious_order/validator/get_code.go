package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/core/rgrouter"
	"member-system-server/internal/app/fictitious_order/api/valid_code"
	"member-system-server/internal/app/fictitious_order/common"
)

/*
 * @Content : validator
 * @Author  : LiJunDong
 * @Time    : 2022-09-14$
 */
type RegisterGetCodeReq struct {
	To string `form:"to" binding:"required" label:"接收方"`
}

func CheckRegisterGetCodeParam(c *gin.Context) {
	this := rgrequest.Get(c)
	var param RegisterGetCodeReq
	err := c.ShouldBind(&param)
	if err != nil {
		errMsg, _ := rgrouter.Error(err)
		this.Response.ReturnError(-500, nil, errMsg)
		return
	}
	if err = param.registerCheckTo(); err != nil {
		this.Response.ReturnError(-500, nil, err.Error())
		return
	}
	if err = param.HighFrequencyRegisterGetCodeLock(this); err != nil {
		this.Response.ReturnError(-500, nil, err.Error())
		return
	}
	this.Param = param
	c.Next()
}

func (m *RegisterGetCodeReq) registerCheckTo() (err error) {
	if rgconfig.GetInt(common.RegisterGetCodeType) == common.SendTypePhone {
		return checkPhone(m.To)
	} else {
		return checkEmail(m.To)
	}
}

// HighFrequencyRegisterGetCodeLock <LiJunDong : 2022-11-06 16:03:57> --- 控制请求频率，获取验证码频率限制
//
//		<LiJunDong : 2023-01-12 10:57:05> --- 简单redis实现方案：
//		用户ID key +1 返回值等于1，过期时间1分钟 如果 +1 后返回值大于2则，返回频繁，更新过期时间为1分钟，如果大于3 则直接返回频繁
//	 <LiJunDong : 2023-01-12 11:04:22> --- 简单mysql实现方案：
//	 查询该账号一分钟之前请求了多少次，大于3次，返回请求频繁
func (m *RegisterGetCodeReq) HighFrequencyRegisterGetCodeLock(this *rgrequest.Client) (err error) {
	client := valid_code.ValidCodeClient{
		This: this,
		To:   m.To,
	}
	err = client.CheckCount()
	return err
}
