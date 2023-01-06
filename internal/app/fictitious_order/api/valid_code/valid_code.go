package valid_code

import (
	"errors"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/util/rgtime"
	"member-system-server/internal/app/fictitious_order/common"
	"member-system-server/pkg/mysql"
	"member-system-server/pkg/mysql/member_system"
	"member-system-server/pkg/random_code"
	"time"
)

/*
 * @Content : valid_code
 * @Author  : LiJunDong
 * @Time    : 2023-01-06$
 */

type ValidCodeClient struct {
	This       *rgrequest.Client
	Code       string
	Msg        string
	Typ        int8
	To         string
	ExpireTime int64
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
	m.Msg = "你的验证码是" + m.Code + "，请勿泄露。"
	m.ExpireTime = rgtime.NowTimeInt() + common.RegisterCodeExpire
	// TODO <LiJunDong : 2023-01-06 18:54:23> --- 开发
	if m.Typ == common.SendTypePhone {

	} else if m.Typ == common.SendTypeEmail {

	}
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
		MsgType:    member_system.ValidCodeMsgTypeRegister,
		Msg:        m.Msg,
		ExpireTime: mysql.Time(expireTime),
	}
	return model.Create(m.This)
}
