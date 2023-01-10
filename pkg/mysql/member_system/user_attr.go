package member_system

import (
	"errors"
	"github.com/jackylee92/rgo/core/rgrequest"
	"gorm.io/gorm"
	"member-system-server/pkg/mysql"
	"strconv"
)

// UserAttr 用户附属表

type UserAttr struct {
	ID               int        `gorm:"primaryKey;column:id;type:int(10);not null"`                  // 主键
	UserID           int        `gorm:"column:user_id;type:int(10);not null;default:0"`              // 用户id
	InvitationCode   string     `gorm:"column:invitation_code;type:varchar(20);not null;default:''"` // 用户推荐码
	InvitationUserID int        `gorm:"column:invitation_user_id;type:int(10);not null;default:0"`   // 推荐的用户id
	Status           int8       `gorm:"column:status;type:tinyint(2);not null;default:0"`            // 状态 0：未知 1：可用 2：禁用
	UpdateTime       mysql.Time `gorm:"column:update_time;type:datetime"`                            // 更新时间
	CreateTime       mysql.Time `gorm:"column:create_time;type:datetime"`                            // 更新时间
	DeleteFlag       int8       `gorm:"column:delete_flag;type:tinyint(2);not null;default:0"`       // 虚拟删除
}

// TableName get sql table name.获取数据库表名
func (m *UserAttr) TableName() string {
	return "user_attr"
}

const (
	UserAttrEffectiveStatus   int8 = 1                       // 有效
	UserAttrUnEffectiveStatus int8 = 1                       // 有效
	UserAttrDefaultStatus     int8 = UserAttrEffectiveStatus // 默认状态
)

var userAttrStatusValue = map[int8]string{
	0: "未知",
	1: "启用",
	2: "禁用",
}

// <LiJunDong : 2023-01-06 19:27:18> --- 好像无效
func (m *UserAttr) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateTime = mysql.NowTime()
	m.UpdateTime = mysql.NowTime()
	m.DeleteFlag = mysql.NoDelete
	return
}

func (m *UserAttr) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (m *UserAttr) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateTime = mysql.NowTime()
	return
}

func (m *UserAttr) Create(this *rgrequest.Client) (err error) {
	model, err := this.Mysql.New("")
	if err != nil {
		return err
	}
	err = model.Db.Table(m.TableName()).Create(m).Error
	return err
}

func (m *UserAttr) Find(param mysql.SearchParam) (exists bool, err error) {
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

func (m *UserAttr) CodeExists(this *rgrequest.Client) (err error) {
	if m.InvitationCode == "" {
		return errors.New("InvitationCode不能为空")
	}
	model := UserAttr{}
	exists, err := model.Find(mysql.SearchParam{
		This:   this,
		Query:  "invitation_code = ? AND status = ?",
		Args:   []interface{}{m.InvitationCode, UserAttrEffectiveStatus},
		Fields: []string{"user_id"},
	})
	if !exists {
		return errors.New("推荐码不存在")
	}
	m.UserID = model.UserID
	return err
}

// TODO <LiJunDong : 2023/1/10 23:56> --- 生成用户自己的推荐码
func CreateUserAttrInvitationCode(userId int) (code string, err error) {
	return "code-" + strconv.Itoa(userId), err
}
