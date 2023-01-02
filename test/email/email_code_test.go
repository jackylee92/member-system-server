package email

import (
	"github.com/jackylee92/rgo"
	"log"
	"member_system-system/pkg/email/email_default"
	"testing"
)

/*
 * @Content : test
 * @Author  : LiJunDong
 * @Time    : 2022-11-06$
 */

/*
* @Content : go test -v -run TestEmailCode test/email/email_code_test.go -count=1 -args -config=../../config/fictitious_order/config.yaml
* @Param   :
* @Return  :
* @Author  : LiJunDong
* @Time    : 2022-05-12
 */
func TestEmailCode(t *testing.T) {
	this := rgo.This
	content := "你好"
	emailClient := email_default.Client{
		This:     this,
		ToEmails: []string{"568915010@qq.com"},
		Content:  content,
	}
	err := emailClient.SendCode()
	log.Println("error: ", err)
}
