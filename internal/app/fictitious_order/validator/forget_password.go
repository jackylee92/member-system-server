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
type ForgetPasswordReq struct {
	Phone string `form:"phone" binding:"required" label:"手机号"`
}

func CheckForgetPasswordParam(c *gin.Context) {
	this := rgrequest.Get(c)
	var param ForgetPasswordReq
	err := c.ShouldBind(&param)
	if err != nil {
		errMsg, _ := rgrouter.Error(err)
		this.Response.ReturnError(-500, nil, errMsg)
		return
	}

	this.Param = param
	c.Next()
}
