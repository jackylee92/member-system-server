package member_system

import (
	"github.com/jackylee92/rgo/core/rgrequest"
	"gorm.io/gorm"
	"log"
	"member-system-server/pkg/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	ID         int        `gorm:"primaryKey;column:id;type:int(10);not null"`
	Username   string     `gorm:"column:username;type:varchar(200);not null"`      // 用户姓名
	CreateTime mysql.Time `gorm:"column:create_time;type:datetime;default:null"`   // 录入时间
	UpdateTime mysql.Time `gorm:"column:update_time;type:datetime;default:null"`   // 更新时间
	DeleteFlag int8       `gorm:"column:delete_flag;type:tinyint(2);default:null"` // 虚拟删除
	Status     int8       `gorm:"column:status;type:tinyint(2);default:null"`
	Remark     string     `gorm:"column:remark;type:varchar(255);default:null"` // 备注
}

// TableName get sql table name.获取数据库表名
func (m *UserInfo) TableName() string {
	return "user_info"
}

var userInfoStatusValue = map[int8]string{
	0: "未知",
	1: "启用",
	2: "禁用",
}

// TODO <LiJunDong : 2023-01-06 19:27:18> --- 好像无效
func (m *UserInfo) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("BeforeCreate----")
	m.CreateTime = mysql.NowTime()
	m.UpdateTime = mysql.NowTime()
	m.DeleteFlag = mysql.NoDelete
	return
}

func (m *UserInfo) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (m *UserInfo) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateTime = mysql.NowTime()
	return
}

func (m *UserInfo) Create(this *rgrequest.Client) (err error) {
	model, err := this.Mysql.New("")
	if err != nil {
		return err
	}
	err = model.Db.Table(m.TableName()).Create(m).Error
	return err
}

func (m *UserInfo) Find(param mysql.SearchParam) (exists bool, err error) {
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
