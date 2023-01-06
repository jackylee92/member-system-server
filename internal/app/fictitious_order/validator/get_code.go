package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/core/rgrouter"
)

/*
 * @Content : validator
 * @Author  : LiJunDong
 * @Time    : 2022-09-14$
 */
type GetCodeReq struct {
	Phone string `form:"phone" binding:"required" label:"手机号"`
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

	this.Param = param
	c.Next()
}

// HighFrequencyRequestLock <LiJunDong : 2022-11-06 16:03:57> --- 未实现 控制请求频率
func HighFrequencyRequestLock(c *gin.Context) {
	c.Next()
}
