package sms

import (
	"github.com/jackylee92/rgo"
	"log"
	"member-system-server/pkg/sms/sms_aliyun"
	"testing"
)

/*
 * @Content : test
 * @Author  : LiJunDong
 * @Time    : 2022-11-06$
 */

/*
* @Content : go test -v -run TestSmsCode test/sms/sms_code_test.go -count=1 -args -config=../../config/fictitious_order/config.yaml
* @Param   :
* @Return  :
* @Author  : LiJunDong
* @Time    : 2022-05-12
 */
func TestSmsCode(t *testing.T) {
	this := rgo.This
	phone := "15755378737"
	code := "1234"
	client := sms_aliyun.Client{
		This:         this,
		Phone:        phone,
		Code:         code,
		Title:        "阿里云短信测试",
		TemplateCode: "SMS_154950909",
	}
	err := client.Send()
	log.Println("error: ", err)
}
