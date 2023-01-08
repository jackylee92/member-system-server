package fictitious_order

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/core/rgrouter"
	"member-system-server/internal/app/fictitious_order/api"
	"member-system-server/internal/app/fictitious_order/message"
	"member-system-server/internal/app/fictitious_order/middleware/app_lock"
	"member-system-server/internal/app/fictitious_order/middleware/authentication"
	"member-system-server/internal/app/fictitious_order/validator"
)

//var f embed.FS

func GetRouter() *gin.Engine {
	registerValidation()
	router := rgrouter.NewRouter()

	publicGroup := router.Group("/public")
	{
		publicGroup.GET("/lock", api.LockHandle)
		publicGroup.POST("/login", validator.CheckLoginParam, api.LoginHandle)
		publicGroup.POST("/register", validator.CheckRegisterParam, api.RegisterHandle)
		publicGroup.GET("/get_code", validator.CheckGetCodeParam, validator.HighFrequencyRequestLock, api.GetCodeHandle)
		// 输入手机号 接受验证码 输入新的密码
		publicGroup.GET("/forget_password", validator.CheckForgetPasswordParam, api.ForgetPasswordHandle)
	}
	homeGroup := router.Group("/home").Use(app_lock.AppLockCheck, authentication.CheckLogin)
	{
		homeGroup.GET("/makeOrder", api.MakeOrderHandle)
	}
	userGroup := router.Group("/user").Use(app_lock.AppLockCheck, authentication.CheckLogin)
	{
		userGroup.GET("/info", validator.CheckGetUserInfoParam, api.GetUserInfoHandle)
		userGroup.GET("/logout", api.LogOutHandle)
	}
	router.NoRoute(func(c *gin.Context) {
		this := rgrequest.Get(c)
		this.Response.ReturnError(message.NoApi)
		return
	})

	return router
}

func registerValidation() {
}
