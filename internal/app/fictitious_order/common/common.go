package common

import (
	"errors"
	"github.com/jackylee92/rgo/core/rgrequest"
)

/*
 * @Content : common
 * @Author  : LiJunDong
 * @Time    : 2022-10-20$
 */

func ReturnErrorAndLog(this *rgrequest.Client, code int64, msg string, err error){
	this.Log.Error("接口返回错误", code, msg, err)
	if msg == "" {
		this.Response.ReturnError(code)
		return
	}
	this.Response.ReturnError(code, nil, msg)
	return
}

func RequestUserId(this *rgrequest.Client)(userId int, err error){
	userIdITF, exists := this.Ctx.Get("user_id")
	if !exists {
		return userId, errors.New("用户ID不存在")
	}
	userId = userIdITF.(int)
	return userId, err
}