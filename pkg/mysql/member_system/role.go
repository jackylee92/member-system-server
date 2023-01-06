package member_system

import (
	"member-system-server/pkg/mysql"
)

// Role 角色表
type Role struct {
	ID         int        `gorm:"primaryKey;unique;column:id;type:int(10);not null"`         // 主键
	Code       string     `gorm:"column:code;type:varchar(50);default:null"`                 // 角色编号
	Name       string     `gorm:"column:name;type:varchar(200);default:null"`                // 角色名称
	Status     int8       `gorm:"column:status;type:tinyint(2);default:null;default:0"`      // 状态
	Remark     string     `gorm:"column:remark;type:varchar(200);default:null"`              // 说明
	CreateTime mysql.Time `gorm:"column:create_time;type:datetime;default:null"`             // 创建时间
	UpdateTime mysql.Time `gorm:"column:update_time;type:datetime;default:null"`             // 更新时间
	DeleteFlag int8       `gorm:"column:delete_flag;type:tinyint(2);default:null;default:0"` // 虚拟删除
	Typ        int8       `gorm:"column:typ;type:tinyint(2);default:null;default:0"`         // 权限类型
}

// TableName get sql table name.获取数据库表名
func (m *Role) TableName() string {
	return "role"
}

var roleStatusValue = map[int8]string{
	0: "未知",
	1: "启用",
	2: "禁用",
}

func (c *Role) Select(param mysql.SearchParam) (list []Role, err error) {
	model, err := param.This.Mysql.New("")
	if err != nil {
		return list, err
	}
	param.Query += " AND delete_flag = 0"
	m := model.Db.Debug().Table(c.TableName()).Where(param.Query, param.Args...)
	if param.Order != "" {
		m.Order(param.Order)
	}
	if param.Limit != 0 {
		m.Limit(param.Limit).Offset(param.Offset)
	}
	if len(param.Fields) != 0 {
		m.Select(param.Fields)
	}
	m.Find(&list)
	return list, err
}
