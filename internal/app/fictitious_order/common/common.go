package common

import (
	"github.com/jackylee92/rgo/core/rgrequest"
)

/*
 * @Content : common
 * @Author  : LiJunDong
 * @Time    : 2022-10-20$
 */

func ReturnErrorAndLog(this *rgrequest.Client, code int64, msg string, err error){
	this.Log.Error("接口返回错误", code, msg, err)
	this.Response.ReturnError(code, nil, msg)
}