package video

import (
	"errors"
	"github.com/jackylee92/rgo/core/rgrequest"
	"gorm.io/gorm"
	"video-admin/pkg/mysql"
)

// ValidCode 验证码记录表
type ValidCode struct {
	ID         int        `gorm:"primaryKey;column:id;type:int(10);not null"`             // 主键
	Code       string     `gorm:"index:index_code;column:code;type:varchar(10);not null"` // 验证码
	Status     int8       `gorm:"column:status;type:tinyint(2);default:null;default:0"`   // 状态 0：未知 1：未使用 2：已使用 3：已失效
	UserID     int        `gorm:"column:user_id;type:int(10);default:null;default:0"`     // 用户ID
	Email      string     `gorm:"column:email;type:varchar(30);default:null"`             // 邮箱
	Phone      string     `gorm:"column:phone;type:varchar(20);default:null"`             // 手机号
	MsgType    int8       `gorm:"column:msg_type;type:tinyint(2);default:null;default:0"` // 消息类型 0：未知 1：注册code
	Msg        string     `gorm:"column:msg;type:varchar(255);default:null"`              // 消息内容
	ExpireTime mysql.Time `gorm:"column:expire_time;type:datetime;default:null"`          // 过期时间
	CreateTime mysql.Time `gorm:"column:create_time;type:datetime;not null"`              // 创建时间
	UpdateTime mysql.Time `gorm:"column:update_time;type:datetime;default:null"`          // 更新时间
	DeleteFlag int8       `gorm:"column:delete_flag;type:tinyint(2);not null;default:0"`  // 虚拟删除 0：未删除 1：已删除
}

// TableName get sql table name.获取数据库表名
func (m *ValidCode) TableName() string {
	return "valid_code"
}

func (m *ValidCode) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateTime = mysql.NowTime()
	m.UpdateTime = mysql.NowTime()
	m.DeleteFlag = 0
	return
}

func (m *ValidCode) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (m *ValidCode) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateTime = mysql.NowTime()
	return
}

func (m *ValidCode) Create(this *rgrequest.Client) (err error) {
	model, err := this.Mysql.New("")
	if err != nil {
		return err
	}
	err = model.Db.Table(m.TableName()).Create(m).Error
	return err
}

func (m *ValidCode) Find(param mysql.SearchParam) (exists bool, err error) {
	model, err := param.This.Mysql.New("")
	if err != nil {
		return exists, err
	}
	param.Query += " AND delete_flag = 0"
	mm := model.Db.Table(m.TableName()).Where(param.Query, param.Args...)
	if param.Fields != nil && len(param.Fields) != 0 {
		mm = mm.Select(param.Fields)
	}
	if param.Order != "" {
		mm = mm.Order(param.Order)
	}
	mm.Find(&m)
	if mm.RowsAffected == 0 {
		return false, err
	}
	return true, mm.Error
}

func (m *ValidCode) Update(param mysql.SearchParam, data map[string]interface{}) (c int64, err error) {
	model, err := param.This.Mysql.New("")
	if err != nil {
		return c, err
	}
	mm := model.Db.Table(m.TableName()).Where(param.Query, param.Args...).Updates(data)
	return mm.RowsAffected, mm.Error
}

func (m *ValidCode) GetCodeByPhone(this *rgrequest.Client) (err error) {
	if len(m.Phone) == 0 {
		return errors.New("phone is nil")
	}
	searchParam := mysql.SearchParam{
		This:   this,
		Query:  "phone = ?",
		Args:   []interface{}{m.Phone},
		Fields: []string{"id", "code", "expire_time", "status"},
		Order:  "id DESC",
	}
	_, err = m.Find(searchParam)
	return err
}

func UseCodeById(this *rgrequest.Client, id int) (err error) {
	model := ValidCode{}
	searchParam := mysql.SearchParam{
		This:  this,
		Query: "id = ?",
		Args:  []interface{}{id},
	}
	data := map[string]interface{}{
		"status": 2,
	}
	c, err := model.Update(searchParam, data)
	if err != nil {
		return err
	}
	if c == 0 {
		return errors.New("验证码数据更新失败")
	}
	return err
}
