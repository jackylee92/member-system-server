package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"member-system-server/internal/app/fictitious_order/api/user"
	"member-system-server/internal/app/fictitious_order/api/valid_code"
	"member-system-server/internal/app/fictitious_order/common"
	"member-system-server/internal/app/fictitious_order/message"
	"member-system-server/internal/app/fictitious_order/validator"
	"member-system-server/pkg/mysql/member_system"
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

type UserListRsp struct {
	List  []UserListRspItem `json:"list"`
	Total int               `json:"total"`
}
type UserListRspItem struct {
	Id              int      `json:"id"`
	Timestamp       int64    `json:"timestamp"`
	Author          string   `json:"author"`
	Reviewer        string   `json:"reviewer"`
	Title           string   `json:"title"`
	ContentShort    string   `json:"content_short"`
	Content         string   `json:"content"`
	Forecast        float64  `json:"forecast"`
	Importance      int      `json:"importance"`
	Type            string   `json:"type"`
	Status          string   `json:"status"`
	DisplayTime     string   `json:"display_time"`
	CommentDisabled bool     `json:"comment_disabled"`
	Pageviews       int      `json:"pageviews"`
	ImageUri        string   `json:"image_uri"`
	Platforms       []string `json:"platforms"`
}

// LoginHandle <LiJunDong : 2023-01-11 11:04:43> --- 登录
func LoginHandle(c *gin.Context) {
	this := rgrequest.Get(c)
	req := this.Param.(validator.LoginReq)
	var msg string
	//  <LiJunDong : 2023-01-11 11:04:58> --- 检查是否已登录
	isLogin, userInfo, token := user.CheckLogin(this)
	defer func() {
		go userInfo.SaveLoginLog(this, req, msg)
	}()
	if isLogin {
		msg = "用户已登录，" + userInfo.Username
		this.Response.ReturnSuccess(getLoginRsp(userInfo, token))
		return
	}
	//  <LiJunDong : 2023-01-11 11:05:29> --- 登录
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

// RegisterGetCodeHandle <LiJunDong : 2023-01-11 11:05:44> --- 注册获取验证码
func RegisterGetCodeHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	req := this.Param.(validator.RegisterGetCodeReq)
	client := valid_code.ValidCodeClient{
		This:  this,
		To:    req.To,
		Typ:   common.SendTypeEmail,
		Scene: member_system.ValidCodeMsgTypeRegister,
	}
	//  <LiJunDong : 2023-01-11 11:06:15> --- 判断获取验证码方式，默认邮箱
	if rgconfig.GetInt(common.RegisterGetCodeType) == common.SendTypePhone {
		client.Typ = common.SendTypePhone
	}
	//  <LiJunDong : 2023-01-11 11:06:36> --- 发送保存验证码
	if err := client.GetCode(); err != nil {
		common.ReturnErrorAndLog(this, -4000, "获取验证码失败", err)
		return
	}
	userInfo := user.Info{
		Account:     req.To,
		ValidCodeId: client.ID,
	}
	//  <LiJunDong : 2023-01-11 11:03:54> --- 生成authorization
	authorization, err := userInfo.GetAuthorization(this, common.JWTTokenRegisterCodeNoUse, common.JWTTokenTypeRegister, client.Code, client.ID)
	if err != nil {
		common.ReturnErrorAndLog(this, -4000, "获取验证码失败", err)
		return
	}
	//  <LiJunDong : 2023-01-11 11:04:14> --- 设置header返回
	this.Response.Ctx.Header("Authorization", authorization)
	this.Response.ReturnSuccess(nil)
	return
}

// LogOutHandle <LiJunDong : 2023-01-11 11:09:07> --- 登出
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

// RegisterHandle <LiJunDong : 2023-01-11 11:09:22> --- 注册
func RegisterHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	defaultUserRolesIds, err := user.DefaultUserRolesIds(this)
	if err != nil {
		this.Log.Error("user.DefaultUserRolesIds", err)
	}
	req := this.Param.(validator.RegisterReq)
	//  <LiJunDong : 2023-01-11 11:09:43> --- 注册
	userInfo := user.Info{
		Account:          req.To,
		Username:         user.DefaultUsername(req.To),
		Password:         req.Password,
		RolesId:          defaultUserRolesIds,
		ValidCodeId:      req.ValidCodeID,
		InvitationUserId: req.InvitationUserId,
	}
	userId, err := userInfo.Register(this)
	if err != nil {
		common.ReturnErrorAndLog(this, -3000, "注册失败【"+err.Error()+"】", err)
		return
	}
	userInfo.UserId = userId
	//  <LiJunDong : 2023-01-11 11:09:57> --- 记录注册日志
	go userInfo.SaveRegisterLog(this, req)
	this.Response.ReturnSuccess(nil)
	return
}

// GetUserInfoHandle <LiJunDong : 2023-01-11 11:10:14> --- 获取用户信息
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

// GetUserListHandle <LiJunDong : 2023-01-11 11:10:27> --- 获取用户列表
func GetUserListHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	param := this.Param.(validator.GetUserListReq)
	this.Log.Info("GetUserListHandle", param)
	userClient := user.ListClient{}
	list, total, err := userClient.GetList()
	if err != nil {
		common.ReturnErrorAndLog(this, message.UserInfoErr, "", err)
	}
	this.Response.ReturnSuccess(getUserListRsp(list, total))
}

func getUserListRsp(list []user.Info, total int) (rsp UserListRsp) {
	rsp.Total = total
	rsp.List = []UserListRspItem{
		{
			Id:              100,
			Timestamp:       1347938534717,
			Author:          "Amy",
			Reviewer:        "Betty",
			Title:           "Uwnbqtpib Fhplkkx Seeyoxm Bmvjuhbmw Lgq Qyydz Wrejbymlnz",
			ContentShort:    "mock data",
			Content:         "<p>I am testing data, I am testing data.</p><p><img src=\"https://wpimg.wallstcn.com/4c69009c-0fd4-4153-b112-6cb53d1cf943\"></p>",
			Forecast:        74.63,
			Importance:      2,
			Type:            "EU",
			Status:          "draft",
			DisplayTime:     "1990-05-13 12:11:24",
			CommentDisabled: true,
			Pageviews:       868,
			ImageUri:        "https://wpimg.wallstcn.com/e4558086-631c-425c-9430-56ffb46e70b3",
			Platforms:       []string{"a-platform"},
		},
		{
			Id:              101,
			Timestamp:       1347938534717,
			Author:          "Amy",
			Reviewer:        "Betty",
			Title:           "Uwnbqtpib Fhplkkx Seeyoxm Bmvjuhbmw Lgq Qyydz Wrejbymlnz",
			ContentShort:    "mock data",
			Content:         "<p>I am testing data, I am testing data.</p><p><img src=\"https://wpimg.wallstcn.com/4c69009c-0fd4-4153-b112-6cb53d1cf943\"></p>",
			Forecast:        74.63,
			Importance:      2,
			Type:            "EU",
			Status:          "draft",
			DisplayTime:     "1990-05-13 12:11:24",
			CommentDisabled: true,
			Pageviews:       868,
			ImageUri:        "https://wpimg.wallstcn.com/e4558086-631c-425c-9430-56ffb46e70b3",
			Platforms:       []string{"a-platform"},
		},
		{
			Id:              102,
			Timestamp:       1347938534717,
			Author:          "Amy",
			Reviewer:        "Betty",
			Title:           "Uwnbqtpib Fhplkkx Seeyoxm Bmvjuhbmw Lgq Qyydz Wrejbymlnz",
			ContentShort:    "mock data",
			Content:         "<p>I am testing data, I am testing data.</p><p><img src=\"https://wpimg.wallstcn.com/4c69009c-0fd4-4153-b112-6cb53d1cf943\"></p>",
			Forecast:        74.63,
			Importance:      2,
			Type:            "EU",
			Status:          "draft",
			DisplayTime:     "1990-05-13 12:11:24",
			CommentDisabled: true,
			Pageviews:       868,
			ImageUri:        "https://wpimg.wallstcn.com/e4558086-631c-425c-9430-56ffb46e70b3",
			Platforms:       []string{"a-platform"},
		},
	}
	return rsp
}

// ForgetGetCodeHandle <LiJunDong : 2023-01-11 11:10:45> --- 忘记密码-获取验证码
func ForgetGetCodeHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	req := this.Param.(validator.ForgetGetCodeReq)
	userInfo := user.Info{
		Account: req.To,
	}
	//  <LiJunDong : 2023-01-11 11:02:54> --- 查询用户账户是否存在
	err := userInfo.FindInfoByAccount(this)
	if err != nil {
		common.ReturnErrorAndLog(this, -4000, err.Error(), err)
		return
	}
	//  <LiJunDong : 2023-01-11 11:03:08> --- 验证码校验
	client := valid_code.ValidCodeClient{
		This:   this,
		To:     req.To,
		Typ:    common.SendTypeEmail,
		Scene:  member_system.ValidCodeMsgTypeForget,
		UserId: userInfo.UserId,
	}
	if rgconfig.GetInt(common.RegisterGetCodeType) == common.SendTypePhone {
		client.Typ = common.SendTypePhone
	}
	if err := client.GetCode(); err != nil {
		common.ReturnErrorAndLog(this, -4000, "获取验证码失败", err)
		return
	}
	//  <LiJunDong : 2023-01-11 11:03:54> --- 生成authorization
	authorization, err := userInfo.GetAuthorization(this, common.JWTTokenForgetCodeNoUse, common.JWTTokenTypeForget, client.Code, client.ID)
	if err != nil {
		common.ReturnErrorAndLog(this, -4000, "获取验证码失败", err)
		return
	}
	//  <LiJunDong : 2023-01-11 11:04:14> --- 设置header返回
	this.Response.Ctx.Header("Authorization", authorization)
	this.Response.ReturnSuccess(nil)
	return
}

// ForgetCheckCodeHandle <LiJunDong : 2023-01-11 11:11:08> --- 检查验证码是否正确
func ForgetCheckCodeHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	req := this.Param.(validator.ForgetCheckCodeReq)
	userInfo := user.Info{
		Account:     req.To,
		ValidCodeId: req.ValidCodeID,
		UserId:      req.UserId,
	}
	//  <LiJunDong : 2023-01-11 11:11:29> --- 检查验证码
	if err := userInfo.ForgetCheckCode(this); err != nil {
		common.ReturnErrorAndLog(this, -4001, "验证失败", err)
		return
	}
	//  <LiJunDong : 2023-01-11 11:11:39> --- 获取authorization
	authorization, err := userInfo.GetAuthorization(this, common.JWTTokenForgetCodeUsed, common.JWTTokenTypeForget, req.ValidCode, req.ValidCodeID)
	if err != nil {
		common.ReturnErrorAndLog(this, -4000, "获取验证码失败", err)
		return
	}
	//  <LiJunDong : 2023-01-11 11:12:02> --- 设置头部
	this.Response.Ctx.Header("Authorization", authorization)
	this.Response.ReturnSuccess(nil)
	return
}

// ForgetNewPasswordHandle <LiJunDong : 2023-01-11 11:12:08> --- 设置新密码
func ForgetNewPasswordHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	req := this.Param.(validator.ForgetNewPasswordReq)
	userInfo := user.Info{
		UserId:   req.UserId,
		Password: req.NewPassword,
	}
	//  <LiJunDong : 2023-01-11 11:12:28> --- 设置新密码
	if err := userInfo.NewPassword(this); err != nil {
		common.ReturnErrorAndLog(this, -4001, "修改密码失败", err)
		return
	}
	go userInfo.SaveNewPasswordLog(this, req)
	this.Response.ReturnSuccess(nil)
	return
}
