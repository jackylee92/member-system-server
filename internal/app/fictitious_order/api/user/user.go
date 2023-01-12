package user

/*
 * @Content : user
 * @Author  : LiJunDong
 * @Time    : 2023-01-05$
 */

import (
	"encoding/json"
	"errors"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"member-system-server/internal/app/fictitious_order/common"
	"member-system-server/pkg/jwt"
	"member-system-server/pkg/mysql"
	"member-system-server/pkg/mysql/member_system"
	"strconv"
	"strings"
)

const (
	defaultStatus = 1 // 默认用户状态
)

type Info struct {
	UserId           int
	Account          string
	Password         string
	Username         string
	Status           int8
	Roles            []string
	RolesId          []int
	Introduction     string
	Avatar           string
	AccountId        int
	ValidCodeId      int
	InvitationCode   string
	InvitationUserId int
}

type ListClient struct {
	This *rgrequest.Client
}

func CheckLogin(this *rgrequest.Client) (res bool, info Info, token string) {
	token = this.Ctx.GetHeader("token")
	if token == "" {
		return false, info, token
	}
	jwtData, err := jwt.ParseToken(this, token)
	if err != nil {
		this.Log.Error("CheckLogin Error", jwtData, err)
		return false, info, token
	}
	userId, err := strconv.Atoi(jwtData.UserId)
	if err != nil {
		this.Log.Error("从JWT中解析用户ID失败", jwtData, err)
		return false, info, token
	}
	info.UserId = userId
	info.Username = jwtData.Username
	return jwtData.Login, info, token
}

func CheckAuthorization(this *rgrequest.Client, typ, status int8) (info Info) {
	token := this.Ctx.GetHeader("Authorization")
	if token == "" {
		return info
	}
	jwtData, err := jwt.ParseToken(this, token)
	if err != nil {
		this.Log.Error("CheckAuthorization Error", jwtData, err)
		return info
	}
	userId, err := strconv.Atoi(jwtData.UserId)
	if err != nil {
		this.Log.Error("从JWT中解析用户ID失败", jwtData, err)
		return info
	}
	jwtDataTyp, _ := strconv.Atoi(jwtData.Typ)
	if jwtDataTyp != int(typ) {
		this.Log.Error("Authorization类型不匹配", jwtData, err)
		return info
	}
	jwtDataStatus, _ := strconv.Atoi(jwtData.Status)
	if int8(jwtDataStatus) != status {
		this.Log.Error("Authorization状态不匹配", jwtData, err)
		return info
	}
	jwtDataValidId, _ := strconv.Atoi(jwtData.ValidId)
	info.UserId = userId
	info.ValidCodeId = jwtDataValidId
	return info
}

func (u *Info) Login(this *rgrequest.Client) (token string, err error) {
	if u.Account == "" || u.Password == "" {
		return token, errors.New("登录失败，用户名密码为空")
	}
	userAccountModel := member_system.UserAccount{Account: u.Account, Password: u.Password}
	err = userAccountModel.GetInfoByAccount(this)
	this.Log.Debug("userAccountModel.GetInfoByAccount", userAccountModel)
	if err != nil {
		this.Log.Error("用户登录失败|" + err.Error())
		return token, errors.New("登录失败，查询账号信息失败")
	}
	if userAccountModel.ID == 0 {
		return token, errors.New("登录失败，用户名或者密码错误")
	}
	if userAccountModel.Status != 1 {
		return token, errors.New("登录失败，账号非启用状态【" + member_system.StatusVal(userAccountModel.Status) + "】")
	}
	userInfoModel := member_system.UserInfo{}
	exists, err := userInfoModel.Find(mysql.SearchParam{
		This:   this,
		Query:  "id = ?",
		Args:   []interface{}{userAccountModel.UserID},
		Fields: []string{"username"},
	})
	if err != nil {
		this.Log.Error("用户登录失败|" + err.Error())
		return token, errors.New("登录失败，查询用户信息失败")
	}
	this.Log.Info("userINfo", exists, userInfoModel)
	if !exists {
		this.Log.Error("用户登录失败|用户信息数据不存在")
		return token, errors.New("登录失败，用户信息数据不存在")
	}
	u.UserId = userAccountModel.UserID
	u.Username = userInfoModel.Username
	jwtData := jwt.LoginData{
		Login:    true,
		UserId:   strconv.Itoa(u.UserId),
		Username: u.Username,
		Typ:      strconv.Itoa(common.JWTTokenTypeLogin),
	}
	token, err = jwt.GetToken(this, jwtData)
	if err != nil {
		this.Log.Error("登录成功，状态保存失败，稍后重试|" + err.Error())
		return token, errors.New("登录成功，状态保存失败，稍后重试")
	}
	return token, err
}

func (u *Info) SaveLoginLog(this *rgrequest.Client, req interface{}, msg string) {
	var logData member_system.UserLog
	logData.UserID = u.UserId
	logData.Type = member_system.UserLogLoginType
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

func (u *Info) SaveLogoutLog(this *rgrequest.Client) {
	var logData member_system.UserLog
	logData.UserID = u.UserId
	logData.Type = member_system.UserLogLogoutType
	logData.Action = "LogoutHandle"
	logData.Remark = "用户退出登录"
	logData.Content = ""
	err := logData.Add(this)
	if err != nil {
		this.Log.Error("保存登出记录失败|" + err.Error())
	}
	return
}

func (u *Info) SaveRegisterLog(this *rgrequest.Client, req interface{}) {
	var logData member_system.UserLog
	logData.UserID = u.UserId
	logData.Type = member_system.UserLogRegisterType
	logData.Action = "RegisterHandle"
	logData.Remark = "注册成功"
	contentByte, _ := json.Marshal(req)
	logData.Content = string(contentByte)
	err := logData.Add(this)
	if err != nil {
		this.Log.Error("保存注册记录失败|" + err.Error())
	}
	return
}

func (m *Info) Register(this *rgrequest.Client) (userId int, err error) {
	if rgconfig.GetBool(common.RegisterCodeOnOffConfig) && m.ValidCodeId == 0 {
		return userId, errors.New("验证码错误")
	}
	accountModel := member_system.UserAccount{
		Account: m.Account,
	}
	exist, err := accountModel.ExistAccount(this)
	if err != nil {
		this.Log.Error(err)
		return userId, errors.New("账号已存在")
	}
	if exist {
		return userId, errors.New("登录名已存在")
	}
	if m.Status == 0 {
		m.Status = defaultStatus
	}
	accountModel.Status = m.Status
	accountModel.Account = m.Account
	accountModel.Password = m.Password

	// <LiJunDong : 2023-01-06 14:00:26> --- 用户信息
	userInfoModel := member_system.UserInfo{
		Username:   m.Username,
		Status:     m.Status,
		DeleteFlag: 0,
	}
	if m.Introduction == "" {
		userInfoModel.Introduction = member_system.UserInfoDefaultIntroduction
	}
	if m.Avatar == "" {
		userInfoModel.Avatar = member_system.UserInfoDefaultAvatar
	}
	err = userInfoModel.Create(this)
	if err != nil {
		this.Log.Error("userInfoModel.Create", err)
		return userId, errors.New("用户信息录入失败")
	}
	accountModel.UserID = userInfoModel.ID
	//  <LiJunDong : 2023-01-06 14:00:40> --- 账号信息
	err = accountModel.Create(this)
	if err != nil {
		this.Log.Error("accountModel.Create", err)
		return userId, errors.New("账号信息录入失败")
	}
	//  <LiJunDong : 2023-01-06 14:00:48> --- 权限信息
	if len(m.RolesId) > 0 {
		userRoleModel := member_system.UserRole{}
		userRoleList := make([]member_system.UserRole, 0, len(m.RolesId))
		for _, item := range m.RolesId {
			userRoleList = append(userRoleList, member_system.UserRole{
				UserID: userInfoModel.ID,
				RoleID: item,
				Status: member_system.UserRoleDefaultStatus,
			})
		}
		err = userRoleModel.BatchCreate(this, userRoleList)
		if err != nil {
			this.Log.Error("userRoleModel.BatchCreate", err)
			return userId, errors.New("权限信息录入失败")
		}
	}
	if rgconfig.GetBool(common.RegisterCodeOnOffConfig) {
		if err = member_system.UseValidCodeById(this, m.ValidCodeId, userInfoModel.ID); err != nil {
			this.Log.Error("member_system.UseValidCodeById", err)
		}
	}
	//  <LiJunDong : 2023/1/11 0:28> --- 创建用户推荐码，记录使用的推荐码
	invitationCode := member_system.CreateUserAttrInvitationCode(userInfoModel.ID)
	//if err != nil {
	//	this.Log.Error("member_system.CreateUserAttrInvitationCode", err)
	//}
	userAttrModel := member_system.UserAttr{
		UserID:           userInfoModel.ID,
		InvitationCode:   invitationCode,
		InvitationUserID: m.InvitationUserId,
		Status:           member_system.UserAttrDefaultStatus,
	}
	err = userAttrModel.Create(this)
	if err != nil {
		this.Log.Error("userAttrModel.Create", err)
	}
	userId = userInfoModel.ID
	return userId, err
}

func DefaultUsername(param string) (nickname string) {
	param = "未命名"
	return param
}

func DefaultUserRolesIds(this *rgrequest.Client) (rolesIds []int, err error) {
	roleModel := member_system.Role{}
	list, err := roleModel.Select(mysql.SearchParam{
		This:   this,
		Query:  "typ = ? AND status = ?",
		Args:   []interface{}{1, 1},
		Fields: []string{"id", "code", "name", "remark"},
	})
	if err != nil {
		return rolesIds, err
	}
	rolesIds = make([]int, 0, len(list))
	for _, item := range list {
		rolesIds = append(rolesIds, item.ID)
	}
	return rolesIds, err
}

func GetInfoById(this *rgrequest.Client, userId int) (userInfo Info, err error) {
	userInfoModel := member_system.UserInfo{}
	exists, err := userInfoModel.Find(mysql.SearchParam{
		This:   this,
		Query:  "id = ?",
		Args:   []interface{}{userId},
		Fields: []string{"id", "username", "introduction", "avatar"},
	})
	if err != nil {
		return userInfo, err
	}
	if !exists {
		return userInfo, errors.New("用户不存在")
	}
	userInfo.UserId = userId
	userInfo.Username = userInfoModel.Username
	userInfo.Introduction = userInfoModel.Introduction
	userInfo.Avatar = userInfoModel.Avatar
	rolesId, roles, err := userInfo.getUserRoles(this)
	if err != nil {
		return userInfo, err
	}
	userInfo.Roles = roles
	userInfo.RolesId = rolesId
	return userInfo, err
}

// getUserRoles 获取权限
func (u *Info) getUserRoles(this *rgrequest.Client) (rolesId []int, roles []string, err error) {
	userRoleModel := member_system.UserRole{}
	userRoleList, err := userRoleModel.Select(mysql.SearchParam{
		This:   this,
		Query:  "user_id = ? AND status = ?",
		Args:   []interface{}{u.UserId, 1},
		Fields: []string{"role_id"},
	})
	if err != nil {
		return rolesId, roles, err
	}
	rolesId = make([]int, 0, len(userRoleList))
	roles = make([]string, 0, len(userRoleList))
	if len(userRoleList) == 0 {
		return rolesId, roles, err
	}
	rolesIdStr := make([]string, 0, len(userRoleList))
	for _, item := range userRoleList {
		rolesId = append(rolesId, item.RoleID)
		rolesIdStr = append(rolesIdStr, strconv.Itoa(item.RoleID))
	}
	roleModel := member_system.Role{}
	roleList, err := roleModel.Select(mysql.SearchParam{
		This:   this,
		Query:  "id IN (" + strings.Join(rolesIdStr, ",") + ") AND status = ?",
		Args:   []interface{}{1},
		Fields: []string{"code"},
	})
	if err != nil {
		return rolesId, roles, err
	}
	for _, item := range roleList {
		roles = append(roles, item.Code)
	}
	return rolesId, roles, err
}

func (c *ListClient) GetList() (list []Info, total int, err error) {
	return list, total, err
}

func (u *Info) ForgetCheckCode(this *rgrequest.Client) (err error) {
	if err := member_system.UseValidCodeById(this, u.ValidCodeId, u.UserId); err != nil {
		this.Log.Error("member_system.UseValidCodeById", err)
		return errors.New("验证码验证失败")
	}
	return err
}

func (u *Info) NewPassword(this *rgrequest.Client) (err error) {
	if u.UserId == 0 {
		return errors.New("用户ID为空")
	}
	if len(u.Password) == 0 {
		return errors.New("用户密码为空")
	}
	accountModel := member_system.UserAccount{}
	err = accountModel.UpdatePassword(mysql.SearchParam{
		This:  this,
		Query: "user_id = ?",
		Args:  []interface{}{u.UserId},
	}, map[string]interface{}{"password": u.Password})
	if err != nil {
		this.Log.Error("accountModel.UpdatePassword", this, err)
	}
	return err
}

func (u *Info) SaveNewPasswordLog(this *rgrequest.Client, req interface{}) {
	var logData member_system.UserLog
	logData.UserID = u.UserId
	logData.Type = member_system.UserLogLoginType
	logData.Action = "NewPasswordHandle"
	logData.Remark = "找回密码成功"
	contentByte, _ := json.Marshal(req)
	logData.Content = string(contentByte)
	err := logData.Add(this)
	if err != nil {
		this.Log.Error("找回密码记录失败|" + err.Error())
	}
	return
}

func (u *Info) FindInfoByAccount(this *rgrequest.Client) (err error) {
	if len(u.Account) == 0 {
		return errors.New("账号不能为空")
	}
	userAccountModel := member_system.UserAccount{}
	exists, err := userAccountModel.Find(mysql.SearchParam{
		This:   this,
		Query:  "account = ?",
		Args:   []interface{}{u.Account},
		Fields: []string{"user_id", "status"},
	})
	if err != nil {
		this.Log.Error("userAccountModel.Find", err)
		return errors.New("获取账户失败")
	}
	if !exists {
		return errors.New("账号不存在")
	}
	if userAccountModel.Status != member_system.UserAccountEffectiveStatus {
		return errors.New("账号不可用")
	}
	u.UserId = userAccountModel.UserID
	return err
}

func (u *Info) GetAuthorization(this *rgrequest.Client, status, typ int8, validCode string, validId int) (authorization string, err error) {
	jwtData := jwt.LoginData{
		Login:     false,
		UserId:    strconv.Itoa(u.UserId),
		Typ:       strconv.Itoa(int(typ)),
		Status:    strconv.Itoa(int(status)),
		ValidCode: validCode,
		ValidId:   strconv.Itoa(validId),
	}
	authorization, err = jwt.GetToken(this, jwtData)
	if err != nil {
		this.Log.Error("jwt.GetToken", jwtData, err)
		return authorization, err
	}
	return authorization, err
}

func (u *Info) CheckInvitationCode(this *rgrequest.Client) (err error) {
	if u.InvitationCode == "" {
		return errors.New("推荐码为空")
	}
	userAttrModel := member_system.UserAttr{
		InvitationCode: u.InvitationCode,
	}
	if err = userAttrModel.CodeExists(this); err != nil {
		return err
	}
	u.InvitationUserId = userAttrModel.UserID
	return err
}
