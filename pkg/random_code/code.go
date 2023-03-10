package random_code

import (
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
)

/*
 * @Content : random_code
 * @Author  : LiJunDong
 * @Time    : 2022-11-06$
 */

func GetCodeInt(min, max int64) (code int64){
	if min > max {
		panic("the min is greater than max!")
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max + 1 + i64Min))
		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

func GetCodeStr(min, max int64)(code string) {
	codeInt := GetCodeInt(min, max)
	code = strconv.Itoa(int(codeInt))
	return code
}