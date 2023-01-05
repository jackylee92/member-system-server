package check_valid_code

import (
	"github.com/jackylee92/rgo/core/rgrequest"
	"github.com/jackylee92/rgo/util/rgtime"
	"member-system-server/pkg/mysql/member_system"
)

/*
 * @Content : check_valid_code
 * @Author  : LiJunDong
 * @Time    : 2022-11-06$
 */

func CheckValidRegisterCode(this *rgrequest.Client, phone, code string) (id int, res bool, err error) {
	model := member_system.ValidCode{Phone: phone}
	err = model.GetCodeByPhone(this)
	if err != nil {
		return id, res, err
	}
	if model.ID == 0 {
		this.Log.Info("验证码不存在", phone, code)
		return id, false, err
	}
	if model.Code != code {
		this.Log.Info("验证码不匹配", phone, model.Code, code)
		return id, false, err
	}
	expireTime := model.ExpireTime.Int()
	if expireTime < rgtime.NowTimeInt() {
		this.Log.Info("验证码已过期", phone, model.Code, code)
		return id, false, err
	}
	return model.ID, true, err
}
