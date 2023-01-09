package member_system

import (
	"errors"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"gorm.io/gorm"
	"member-system-server/internal/app/fictitious_order/common"
	"member-system-server/pkg/mysql"
	"time"
)

// UserAccount [...]
type UserAccount struct {
	ID         int        `gorm:"primaryKey;column:id;type:int(10);not null"`
	UserID     int        `gorm:"column:user_id;type:int(10);default:null;default:0"`    // 用户ID
	Account    string     `gorm:"column:account;type:varchar(100);not null;default:''"`  // 登录名
	Password   string     `gorm:"column:password;type:varchar(150);not null;default:''"` // 密码
	CreateTime mysql.Time `gorm:"column:create_time;type:datetime;not null"`             // 创建时间
	UpdateTime mysql.Time `gorm:"column:update_time;type:datetime;not null"`             // 更新时间
	DeleteFlag int8       `gorm:"column:delete_flag;type:tinyint(2);not null;default:0"` // 虚拟删除 0:未删除 1:已删除
	Status     int8       `gorm:"column:status;type:tinyint(2);not null;default:0"`      // 状态0:未知 1:启用 2:禁用
}

// TableName get sql table name.获取数据库表名
func (m *UserAccount) TableName() string {
	return "user_account"
}

var userAccountStatusValue = map[int8]string{
	0: "未知",
	1: "启用",
	2: "禁用",
}

func (m *UserAccount) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateTime = mysql.NowTime()
	m.UpdateTime = mysql.NowTime()
	m.DeleteFlag = 0
	return
}

func (m *UserAccount) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (m *UserAccount) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateTime = mysql.NowTime()
	return
}

func (m *UserAccount) Find(param mysql.SearchParam) (exists bool, err error) {
	model, err := param.This.Mysql.New("")
	if err != nil {
		return exists, err
	}
	param.Query += " AND delete_flag = 0"
	mm := model.Db.Table(m.TableName()).Debug().Where(param.Query, param.Args...)
	if param.Fields != nil && len(param.Fields) != 0 {
		mm = mm.Select(param.Fields)
	}
	mm.Find(&m)
	if mm.RowsAffected == 0 {
		return false, err
	}
	return true, mm.Error
}

func (m *UserAccount) Create(this *rgrequest.Client) (err error) {
	m.Password = getPassword(m.Password)
	model, err := this.Mysql.New("")
	if err != nil {
		return err
	}
	err = model.Db.Table(m.TableName()).Create(m).Error
	return err
}

func (m *UserAccount) GetInfoByAccount(this *rgrequest.Client) (err error) {
	searchParam := mysql.SearchParam{
		Query:  "account = ? AND password = ?",
		Args:   []interface{}{m.Account, getPassword(m.Password)},
		Fields: []string{"id", "status", "account", "user_id"},
		This:   this,
	}
	_, err = m.Find(searchParam)
	if err != nil {
		return errors.New("通过账号查询用户信息失败|" + err.Error())
	}
	return err
}

// TODO <LiJunDong : 2022-11-04 18:35:54> --- 加密
func getPassword(password string) (newPassword string) {
	return password + rgconfig.GetStr(common.UserPasswordSaltKey)
}

func StatusVal(status int8) string {
	value, ok := userAccountStatusValue[status]
	if !ok {
		value = userAccountStatusValue[0]
	}
	return value
}

func (m *UserAccount) ExistAccount(this *rgrequest.Client) (exist bool, err error) {
	if len(m.Account) == 0 {
		return false, errors.New("account为空")
	}
	model := UserAccount{}
	searchParam := mysql.SearchParam{
		Query:  "account = ?",
		Args:   []interface{}{m.Account},
		This:   this,
		Fields: []string{"id"},
	}
	return model.Find(searchParam)
}

func (m *UserAccount) UpdatePassword(param mysql.SearchParam, data map[string]interface{}) (err error) {
	password, ok := data["password"]
	if !ok {
		return errors.New("密码不能为空")
	}
	if len(password.(string)) == 0 {
		return errors.New("密码不能为空")
	}
	data["update_time"] = time.Now()
	model, err := param.This.Mysql.New("")
	if err != nil {
		return err
	}
	return model.Db.Table(m.TableName()).Where(param.Query, param.Args...).Updates(data).Error
}
