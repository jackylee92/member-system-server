package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgrequest"
	"member-system-server/internal/app/fictitious_order/api/user"
	"member-system-server/internal/app/fictitious_order/common"
	"member-system-server/internal/app/fictitious_order/message"
	"member-system-server/internal/app/fictitious_order/validator"
)

type LoginRsp struct {
	Token string `json:"token"`
}

type UserInfoRsp struct {
	Username     string   `json:"username"`
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	UserId       int      `json:"user_id"`
}

func LoginHandle(c *gin.Context) {
	this := rgrequest.Get(c)
	req := this.Param.(validator.LoginReq)
	var msg string
	isLogin, userInfo, token := user.CheckLogin(this)
	defer func() {
		go userInfo.SaveLoginLog(this, req, msg)
	}()
	if isLogin {
		msg = "用户已登录，" + userInfo.Username
		this.Response.ReturnSuccess(getLoginRsp(userInfo, token))
		return
	}
	userInfo.Account = req.Account
	userInfo.Password = req.Password
	token, err := userInfo.Login(this)
	this.Log.Debug("login", userInfo, err)
	if err != nil {
		msg = err.Error()
		this.Response.ReturnError(-1000, nil, msg)
		return
	}
	this.Log.Info("userInfo", userInfo)
	msg = "用户登录，登录成功，" + userInfo.Username
	this.Response.ReturnSuccess(getLoginRsp(userInfo, token))
	return
}

func getLoginRsp(userInfo user.Info, token string) (rsq LoginRsp) {
	rsq.Token = token
	return rsq
}

func LogOutHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	userInfo := user.Info{}
	isLogin, userInfo, _ := user.CheckLogin(this)
	if !isLogin {
		return
	}
	userInfo.SaveLogoutLog(this)
	this.Response.ReturnSuccess(nil)
}

func RegisterHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	defaultUserRolesIds, err := user.DefaultUserRolesIds(this)
	if err != nil {
		this.Log.Error("user.DefaultUserRolesIds", err)
	}
	req := this.Param.(validator.RegisterReq)
	userInfo := user.Info{
		Account:     req.Phone,
		Username:    user.DefaultUsername(req.Phone),
		Password:    req.Password,
		RolesId:     defaultUserRolesIds,
		ValidCodeId: req.ValidCodeID,
	}
	userId, err := userInfo.Register(this)
	if err != nil {
		common.ReturnErrorAndLog(this, -3000, "注册失败【"+err.Error()+"】", err)
		return
	}
	userInfo.UserId = userId
	go userInfo.SaveRegisterLog(this, req)
	this.Response.ReturnSuccess(nil)
	return
}

func ForgetPasswordHandle(ctx *gin.Context) {
	return
}

func GetUserInfoHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	userId, err := common.RequestUserId(this)
	if err != nil {
		common.ReturnErrorAndLog(this, message.NoUserId, "", err)
		return
	}
	userInfo, err := user.GetInfoById(this, userId)
	if err != nil {
		common.ReturnErrorAndLog(this, message.UserInfoErr, "", err)
	}
	this.Response.ReturnSuccess(getUserInfoRsp(userInfo))
}

func getUserInfoRsp(userInfo user.Info) (rsp UserInfoRsp) {
	rsp.UserId = userInfo.UserId
	rsp.Username = userInfo.Username
	rsp.Roles = userInfo.Roles
	rsp.Avatar = userInfo.Avatar
	rsp.Introduction = userInfo.Introduction
	return rsp
}
