package msg_count

import (
	"github.com/jackylee92/rgo"
	_ "github.com/jackylee92/rgo"
	"log"
	"member-system-server/pkg/mysql/member_system"
	"testing"
)

/*
 * @Content : test
 * @Author  : LiJunDong
 * @Time    : 2022-11-06$
 */

/*
* @Content : go test -v -run TestMsgCount test/msg_count/msg_count_test.go -count=1 -args -config=../../config/fictitious_order/config.yaml
* @Param   :
* @Return  :
* @Author  : LiJunDong
* @Time    : 2022-05-12
 */
func TestMsgCount(t *testing.T) {
	this := rgo.This
	model := member_system.ValidCode{}
	c, err := model.GetLastMinuteCount(this, "jackylee92@139.com")
	log.Println(c, err)
}
