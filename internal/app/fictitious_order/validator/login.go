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
type LoginReq struct {
	Account  string `json:"account" binding:"required" label:"账号"`
	Password string `json:"password" binding:"required" label:"密码"`
}

func CheckLoginParam(c *gin.Context) {
	this := rgrequest.Get(c)
	var param LoginReq
	err := c.ShouldBind(&param)
	if err != nil {
		errMsg, _ := rgrouter.Error(err)
		this.Response.ReturnError(-500, nil, errMsg)
		return
	}

	this.Param = param
	c.Next()
}
