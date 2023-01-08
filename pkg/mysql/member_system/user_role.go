package member_system

import (
	"github.com/jackylee92/rgo/core/rgrequest"
	"gorm.io/gorm"
	"member-system-server/pkg/mysql"
	"strconv"
)

// UserRole 用户角色关系
type UserRole struct {
	ID         int        `gorm:"primaryKey;unique;column:id;type:int(10);not null;default:0"` // 主键ID
	UserID     int        `gorm:"column:user_id;type:int(10);default:null;default:0"`          // 用户ID
	RoleID     int        `gorm:"column:role_id;type:int(10);default:null;default:0"`          // 角色ID
	UpdateTime mysql.Time `gorm:"column:update_time;type:datetime;default:null"`               // 更新时间
	Status     int8       `gorm:"column:status;type:tinyint(2);default:null;default:0"`        // 状态
	DeleteFlag int8       `gorm:"column:delete_flag;type:tinyint(2);default:null;default:0"`   // 虚拟删除
	Remark     string     `gorm:"column:remark;type:varchar(200);default:null"`                // 备注
	CreateTime mysql.Time `gorm:"column:create_time;type:datetime;default:null"`               // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *UserRole) TableName() string {
	return "user_role"
}

var userRoleStatusValue = map[int8]string{
	0: "未知",
	1: "启用",
	2: "禁用",
}

var UserRoleDefaultStatus int8 = 1

func (m *UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateTime = mysql.NowTime()
	m.UpdateTime = mysql.NowTime()
	m.DeleteFlag = 0
	return
}

func (m *UserRole) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (m *UserRole) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateTime = mysql.NowTime()
	return
}

func (m *UserRole) Create(this *rgrequest.Client) (err error) {
	model, err := this.Mysql.New("")
	if err != nil {
		return err
	}
	err = model.Db.Table(m.TableName()).Create(m).Error
	return err
}

func (m *UserRole) BatchCreate(this *rgrequest.Client, param []UserRole) (err error) {
	model, err := this.Mysql.New("")
	if err != nil {
		return err
	}
	return model.Db.Table(m.TableName()).CreateInBatches(param, 50).Error
}

func (m *UserRole) Find(param mysql.SearchParam) (exists bool, err error) {
	model, err := param.This.Mysql.New("")
	if err != nil {
		return exists, err
	}
	param.Query += " AND delete_flag = " + strconv.Itoa(int(mysql.NoDelete))
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

func (m *UserRole) Select(param mysql.SearchParam) (list []UserRole, err error) {
	model, err := param.This.Mysql.New("")
	if err != nil {
		return list, err
	}
	param.Query += " AND delete_flag = " + strconv.Itoa(int(mysql.NoDelete))
	mm := model.Db.Table(m.TableName()).Where(param.Query, param.Args...)
	if param.Order != "" {
		mm.Order(param.Order)
	}
	if param.Limit != 0 {
		mm.Limit(param.Limit).Offset(param.Offset)
	}
	if len(param.Fields) != 0 {
		mm.Select(param.Fields)
	}
	mm.Find(&list)
	return list, err
}
