package app_lock

import (
	"github.com/gin-gonic/gin"
)

/*
 * @Content : app_lock
 * @Author  : LiJunDong
 * @Time    : 2022-09-13$
 */

func AppLockCheck(c *gin.Context) {
	//if !home.AppLock {
	//	this := rgrequest.Get(c)
	//	this.Response.ReturnError(message.AppLockCode)
	//	return
	//}
	c.Next()
}
