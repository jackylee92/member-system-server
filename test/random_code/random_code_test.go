package email

import (
	_ "github.com/jackylee92/rgo"
	"log"
	"member-system-server/pkg/random_code"
	"testing"
)

/*
 * @Content : test
 * @Author  : LiJunDong
 * @Time    : 2022-11-06$
 */

/*
* @Content : go test -v -run TestRandomCode test/random_code/random_code_test.go -count=1 -args -config=../../config/fictitious_order/config.yaml
* @Param   :
* @Return  :
* @Author  : LiJunDong
* @Time    : 2022-05-12
 */
func TestRandomCode(t *testing.T) {
	code := random_code.GetCodeInt(1000, 9999)
	log.Println("code: ", code)
}
