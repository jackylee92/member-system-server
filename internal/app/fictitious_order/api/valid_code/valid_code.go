package valid_code

import (
	"errors"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/util/rgtime"
	"member-system-server/internal/app/fictitious_order/common"
	"member-system-server/pkg/email/email_default"
	"member-system-server/pkg/mysql"
	"member-system-server/pkg/mysql/member_system"
	"member-system-server/pkg/random_code"
	"member-system-server/pkg/sms/sms_aliyun"
	"time"
)

/*
 * @Content : valid_code
 * @Author  : LiJunDong
 * @Time    : 2023-01-06$
 */

type ValidCodeClient struct {
	This       *rgrequest.Client
	ID         int
	Code       string
	Msg        string
	Typ        int8
	To         string
	ExpireTime int64
	Template   int    // 模版ID
	Scene      int8   // 场景
	DeviceInfo string // 设备信息，记录这个code指定某个设备可用
	UserId     int    // 用户id
}

func (m *ValidCodeClient) GetCode() (err error) {
	if err := m.setCode(); err != nil {
		m.This.Log.Error("m.setCode", err)
		return errors.New("获取验证码失败")
	}
	if err := m.send(); err != nil {
		m.This.Log.Error("m.send", err)
		return errors.New("发送验证码失败")
	}
	if err := m.save(); err != nil {
		m.This.Log.Error("m.save", err)
		return errors.New("保存验证码失败")
	}
	m.This.Log.Info("验证码", m.To, m.Code)
	return
}

func (m *ValidCodeClient) setCode() (err error) {
	code := random_code.GetCodeStr(common.RegisterCodeMin, common.RegisterCodeMax)
	m.Code = code
	return err
}

func (m *ValidCodeClient) send() (err error) {
	if err = m.setContent(); err != nil {
		return err
	}
	m.ExpireTime = rgtime.NowTimeInt() + common.RegisterCodeExpire
	if m.Typ == common.SendTypePhone {
		client := sms_aliyun.Client{
			This:         m.This,
			Phone:        m.To,
			Code:         m.Code,
			Title:        "阿里云短信测试",
			TemplateCode: "SMS_154950909",
		}
		err = client.Send()
	} else if m.Typ == common.SendTypeEmail {
		client := email_default.Client{
			This:     m.This,
			Content:  m.Msg,
			ToEmails: []string{m.To},
		}
		err = client.Send()
	}
	if err != nil {
		m.This.Log.Error("验证码发送失败", err, m)
	}
	return err
}

func (m *ValidCodeClient) setContent() (err error) {
	m.Msg = "你的验证码是" + m.Code + "，请勿泄露。"
	return err
}

func (m *ValidCodeClient) save() (err error) {
	expireTime := time.Unix(m.ExpireTime, 64)
	var phone, email string
	if m.Typ == common.SendTypePhone {
		phone = m.To
	} else if m.Typ == common.SendTypeEmail {
		email = m.To
	}
	model := member_system.ValidCode{
		Code:       m.Code,
		Phone:      phone,
		Email:      email,
		Status:     member_system.ValidCodeDefaultStatus,
		MsgType:    m.Scene,
		Msg:        m.Msg,
		ExpireTime: mysql.Time(expireTime),
		DeviceInfo: m.DeviceInfo,
		UserID:     m.UserId,
	}
	err = model.Create(m.This)
	m.ID = model.ID
	return err
}
