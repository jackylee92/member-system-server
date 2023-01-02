package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
)
var AppLock bool

func LockHandle(c *gin.Context){
	this := rgrequest.Get(c)
	res := this.Ctx.Query("lock")
	msg := ""
	if res == "ljd" {
		AppLock = true
		msg = "已开启"
	}else{
		AppLock = false
		msg = "已关闭"
	}
	this.Response.ReturnSuccess(msg)
	return
}
