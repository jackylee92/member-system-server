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
)

const (
	defaultStatus = 1 // 默认用户状态
)

type Info struct {
	UserId       int
	Account      string
	Password     string
	Username     string
	Status       int8
	Roles        []string
	RolesId      []int
	Introduction string
	Avatar       string
	AccountId    int
	ValidCodeId  int
}

func CheckLogin(this *rgrequest.Client) (res bool, info Info, token string) {
	token = this.Ctx.Query("token")
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
		if err = member_system.UseValidCodeById(this, m.ValidCodeId); err != nil {
			this.Log.Error("member_system.UseValidCodeById", err)
		}
	}
	userId = userInfoModel.ID
	return userId, err
}

// TODO <LiJunDong : 2022-11-04 18:36:14> --- 开发
func DefaultUsername(param string) (nickname string) {
	param = "username-" + param
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
		Fields: []string{"id", "username"},
	})
	if err != nil {
		return userInfo, err
	}
	if !exists {
		return userInfo, errors.New("用户不存在")
	}
	userInfo.UserId = userId
	userInfo.Username = userInfoModel.Username
	return userInfo, err
}
