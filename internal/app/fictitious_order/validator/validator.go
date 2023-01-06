package validator

import (
	"errors"
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/util/rgtime"
	"member-system-server/internal/app/fictitious_order/common"
	"member-system-server/pkg/mysql/member_system"
)

/*
 * @Content : validator
 * @Author  : LiJunDong
 * @Time    : 2023-01-06$
 */

// CheckValidCode <LiJunDong : 2023-01-06 13:56:53> --- 验证验证码
func CheckValidCode(this *rgrequest.Client, typ int8, to, code string) (id int, err error) {
	model := member_system.ValidCode{Phone: to}
	if typ == common.SendTypeEmail {
		model.Email = to
		err = model.GetCodeByEmail(this)
	} else if typ == common.SendTypePhone {
		model.Phone = to
		err = model.GetCodeByPhone(this)
	} else {
		return id, errors.New("发送方式错误")
	}
	if err != nil {
		return id, err
	}
	if model.ID == 0 {
		this.Log.Debug("model.GetCodeByPhone", to, code)
		return id, errors.New("验证码不存在")
	}
	if model.Code != code {
		this.Log.Debug("model.Code != code", to, model.Code, code)
		return id, errors.New("验证码不匹配")
	}
	if model.Status != member_system.UsableValidCodeStatus {
		this.Log.Debug("model.Code != code", to, model.Code, code)
		return id, errors.New("验证码不匹配")
	}
	expireTime := model.ExpireTime.Int()
	if expireTime < rgtime.NowTimeInt() {
		this.Log.Debug("expireTime < rgtime.NowTimeInt()", to, model.Code, code)
		return id, errors.New("验证码已过期")
	}
	return model.ID, err
}
