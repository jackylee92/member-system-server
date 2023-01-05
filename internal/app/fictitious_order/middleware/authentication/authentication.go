package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"member-system-server/internal/app/fictitious_order/api"
	"member-system-server/internal/app/fictitious_order/message"
)

/*
 * @Content : authentication
 * @Author  : LiJunDong
 * @Time    : 2022-09-18$
 */

const (
	jwtHeaderName = "jwt_header_name"
)

func CheckLogin(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	isLogin, _, _ := api.CheckLogin(this)
	if !isLogin {
		this.Response.ReturnError(message.NoLoginCode)
		return
	}
	this.Ctx.Next()
}
