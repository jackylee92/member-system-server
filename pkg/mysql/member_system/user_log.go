package member_system

import (
	"github.com/jackylee92/rgo/core/rgrequest"
	"gorm.io/gorm"
	"member-system-server/pkg/mysql"
)

// UserLog [...]
type UserLog struct {
	ID         int        `gorm:"primaryKey;column:id;type:int(11);not null"`
	UserID     int        `gorm:"column:user_id;type:int(11);not null;default:0"`        // 用户ID
	Type       int8       `gorm:"column:type;type:tinyint(255);not null;default:0"`      // 类型 0:未知 1:登录 2:登出
	Content    string     `gorm:"column:content;type:text;not null"`                     // 提交数据
	Action     string     `gorm:"column:action;type:varchar(30);not null;default:''"`    // 操作
	CreateTime mysql.Time `gorm:"column:create_time;type:datetime;not null"`             // 创建时间
	UpdateTime mysql.Time `gorm:"column:update_time;type:datetime;not null"`             // 更新时间
	Status     int8       `gorm:"column:status;type:tinyint(2);not null;default:0"`      // 状态 0:未知 1:正常
	DeleteFlag int8       `gorm:"column:delete_flag;type:tinyint(2);not null;default:0"` // 虚拟删除 0:未删除 1:已删除
	Remark     string     `gorm:"column:remark;type:varchar(255);not null;default:''"`   // 一句话说明
}

// TableName get sql table name.获取数据库表名
func (m *UserLog) TableName() string {
	return "user_log"
}

var UserLogLoginType int8 = 1    // 登录类型
var UserLogLogoutType int8 = 2   // 退出类型
var UserLogRegisterType int8 = 3 // 注册类型

func (m *UserLog) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateTime = mysql.NowTime()
	m.UpdateTime = mysql.NowTime()
	m.DeleteFlag = 0
	return
}

func (m *UserLog) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (m *UserLog) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateTime = mysql.NowTime()
	return
}

func (m *UserLog) Add(this *rgrequest.Client) (err error) {
	model, err := this.Mysql.New("")
	if err != nil {
		return err
	}
	m.Status = 1
	mm := model.Db.Table(m.TableName()).Create(m)
	err = mm.Error
	if err != nil {
		return err
	}
	return err
}
