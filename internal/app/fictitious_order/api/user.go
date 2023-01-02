package api

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"video-admin/internal/app/fictitious_order/common"
	"video-admin/internal/app/fictitious_order/validator"
	"video-admin/pkg/jwt"
	"video-admin/pkg/mysql/video"
)

const (
	sessionIsLoginKey = "is_login"
	sessionUserInfo   = "userinfo"
	defaultUserStatus = 1
)

func init() {
	gob.Register(UserInfo{})
}

type UserInfo struct {
	Username     string
	Password     string
	Status       int8
	Nickname     string
	Roles        []string
	Introduction string
	Avatar       string
	Id           int
}

type LoginRsp struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	Token    string `json:"token"`
}

type UserInfoRsp struct {
	Username     string `json:"username"`
	Nickname     string `json:"name"`
	Roles        []string `json:"roles"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
	Id           int `json:"id"`
}

func CheckLogin(this *rgrequest.Client) (res bool, userInfo UserInfo, token string) {
	session := sessions.Default(this.Ctx)
	isLoginITF := session.Get(sessionIsLoginKey)
	this.Log.Info("isLogin", isLoginITF)
	if isLoginITF == nil {
		return false, userInfo, token
	}
	isLogin, ok := isLoginITF.(bool)
	if !ok {
		return false, userInfo, token
	}
	if !isLogin {
		return false, userInfo, token
	}
	userInfoITF := session.Get(sessionUserInfo)
	if userInfoITF == nil {
		return false, userInfo, token
	}
	userInfo, ok = userInfoITF.(UserInfo)
	if !ok {
		return false, userInfo, token
	}
	token = session.ID()
	return true, userInfo, token
}

func LoginHandle(c *gin.Context) {
	this := rgrequest.Get(c)
	req := this.Param.(validator.LoginReq)
	var msg string
	isLogin, userInfo, token := CheckLogin(this)
	defer func() {
		go userInfo.saveLoginLog(this, req, msg)
	}()
	if isLogin {
		msg = "用户已登录，" + userInfo.Username
		this.Response.ReturnSuccess(getLoginRsp(userInfo, token))
		return
	}
	userInfo.Username = req.Username
	userInfo.Password = req.Password
	token, err := userInfo.login(this)
	this.Log.Debug("login", userInfo, err)
	if err != nil {
		msg = err.Error()
		this.Response.ReturnError(-1000, nil, msg)
		return
	}
	msg = "用户登录，登录成功，" + userInfo.Username
	this.Response.ReturnSuccess(getLoginRsp(userInfo, token))
	return
}

func getLoginRsp(req UserInfo, token string) (rsq LoginRsp) {
	rsq.Username = req.Username
	rsq.Id = req.Id
	rsq.Token = token
	return rsq
}

func (u *UserInfo) login(this *rgrequest.Client) (token string, err error) {
	if u.Username == "" || u.Password == "" {
		return token, errors.New("登录失败，用户名密码为空")
	}
	userAccountModel := video.UserAccount{Username: u.Username, Password: u.Password}
	err = userAccountModel.GetUserInfoByAccount(this)
	this.Log.Debug("userAccountModel.GetUserInfoByAccount", userAccountModel)
	if err != nil {
		this.Log.Error("用户登录失败|" + err.Error())
		return token, errors.New("登录失败，查询失败")
	}
	if userAccountModel.ID == 0 {
		return token, errors.New("登录失败，用户名或者密码错误")
	}
	if userAccountModel.Status != 1 {
		return token, errors.New("登录失败，非启用状态【" + video.StatusVal(userAccountModel.Status) + "】")
	}
	u.Id = userAccountModel.ID
	jwtData := jwt.LoginData{
		Login: true,
		UserId: u.Id,
		Username: u.Username,
	}
	token, err = jwt.GetToken(this, jwtData)
	if err != nil {
		this.Log.Error("登录成功，状态保存失败，稍后重试|" + err.Error())
		return token, errors.New("登录成功，状态保存失败，稍后重试")
	}
	return token, err
}

func (u *UserInfo) saveLoginLog(this *rgrequest.Client, req validator.LoginReq, msg string) {
	var logData video.UserLog
	logData.UserID = u.Id
	logData.Type = 1
	logData.Action = "LoginHandle"
	logData.Remark = msg
	contentByte, _ := json.Marshal(req)
	logData.Content = string(contentByte)
	err := logData.Add(this)
	if err != nil {
		this.Log.Error("保存登录记录失败|" + err.Error())
	}
	return
}

func LogOutHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	session := sessions.Default(this.Ctx)
	userInfoITF := session.Get(sessionUserInfo)
	if userInfoITF == nil {
		this.Response.ReturnError(-2000)
		return
	}
	userInfo, ok := userInfoITF.(UserInfo)
	if !ok {
		this.Response.ReturnError(-2001)
		return
	}
	session.Clear()
	userInfo.saveLogoutLog(this)
	this.Response.ReturnSuccess(nil)
}

func (u *UserInfo) saveLogoutLog(this *rgrequest.Client) {
	var logData video.UserLog
	logData.UserID = u.Id
	logData.Type = 2
	logData.Action = "LogoutHandle"
	logData.Remark = "用户退出登录"
	logData.Content = ""
	err := logData.Add(this)
	if err != nil {
		this.Log.Error("保存登录记录失败|" + err.Error())
	}
	return
}

func RegisterHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	req := this.Param.(validator.RegisterReq)
	userInfo := UserInfo{
		Username: req.Phone,
		Nickname: defaultNickname(req.Phone),
		Password: req.Password,
		Status:   1,
	}
	err := userInfo.Register(this)
	if err != nil {
		common.ReturnErrorAndLog(this, -3000, "注册失败", err)
		return
	}
	if rgconfig.GetBool(common.RegisterCodeOnOffConfig) {
		go video.UseCodeById(this, req.ValidCodeID)
	}
	this.Response.ReturnSuccess(nil)
	return
}

func (m *UserInfo) Register(this *rgrequest.Client) (err error) {
	model := video.UserAccount{
		Username: m.Username,
	}
	exist, err := model.ExistUsername(this)
	if err != nil {
		this.Log.Error(err)
		return errors.New("注册失败")
	}
	if exist {
		return errors.New("用户名已存在")

	}
	model.Status = m.Status
	model.Username = m.Username
	model.Nickname = m.Nickname
	model.Password = m.Password
	err = model.Create(this)
	return err
}

// TODO <LiJunDong : 2022-11-04 18:36:14> --- 开发
func defaultNickname(phone string) (nickname string) {
	return phone
}

func ForgetPasswordHandle(ctx *gin.Context) {
	return
}

func GetUserInfoHandle(ctx *gin.Context) {
	this := rgrequest.Get(ctx)
	userInfo := UserInfo{
		Id: 1,
		Nickname: "ljd",
		Roles: []string{"admin"},
		Introduction: "接口返回",
		Avatar: "",
	}
	this.Response.ReturnSuccess(getUserInfoRsp(userInfo))
}

func getUserInfoRsp(userInfo UserInfo) (rsp UserInfoRsp) {
	rsp.Id = userInfo.Id
	rsp.Username = userInfo.Username
	rsp.Roles = userInfo.Roles
	rsp.Nickname = userInfo.Nickname
	rsp.Avatar = userInfo.Avatar
	rsp.Introduction = userInfo.Introduction
	return rsp
}
