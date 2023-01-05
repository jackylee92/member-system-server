package member_system

import (
	"errors"
	"github.com/jackylee92/rgo/core/rgconfig"
	"github.com/jackylee92/rgo/core/rgrequest"
	"gorm.io/gorm"
	"member-system-server/internal/app/fictitious_order/common"
	"member-system-server/pkg/mysql"
)

// UserAccount [...]
type UserAccount struct {
	ID         int        `gorm:"primaryKey;column:id;type:int(10);not null"`
	Username   string     `gorm:"column:username;type:varchar(100);not null;default:''"`
	Password   string     `gorm:"column:password;type:varchar(150);not null;default:''"`
	Status     int8       `gorm:"column:status;type:tinyint(2);not null;default:0"`
	CreateTime mysql.Time `gorm:"column:create_time;type:datetime;not null"`
	UpdateTime mysql.Time `gorm:"column:update_time;type:datetime;not null"`
	Nickname   string     `gorm:"column:nickname;type:varchar(255);not null;default:''"`
	DeleteFlag int8       `gorm:"column:delete_flag;type:tinyint(2);not null;default:0"`
}

var statusValue = map[int8]string{
	0: "未知",
	1: "启用",
	2: "禁用",
}

// TableName get sql table name.获取数据库表名
func (m *UserAccount) TableName() string {
	return "user_account"
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
	mm := model.Db.Table(m.TableName()).Where(param.Query, param.Args...)
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
	model, err := this.Mysql.New("")
	if err != nil {
		return err
	}
	err = model.Db.Table(m.TableName()).Create(m).Error
	return err
}

func (m *UserAccount) GetUserInfoByAccount(this *rgrequest.Client) (err error) {
	searchParam := mysql.SearchParam{
		Query:  "username = ? AND password = ?",
		Args:   []interface{}{m.Username, getPassword(m.Password)},
		Fields: []string{"id", "status", "nickname"},
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
	value, ok := statusValue[status]
	if !ok {
		value = statusValue[0]
	}
	return value
}

func (m *UserAccount) ExistUsername(this *rgrequest.Client) (exist bool, err error) {
	if len(m.Username) == 0 {
		return false, errors.New("username为空")
	}
	model := UserAccount{}
	searchParam := mysql.SearchParam{
		Query:  "username = ?",
		Args:   []interface{}{m.Username},
		This:   this,
		Fields: []string{"id"},
	}
	return model.Find(searchParam)
}
