package mysql

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/jackylee92/rgo/core/rgglobal/rgconst"
	"github.com/jackylee92/rgo/core/rgrequest"
	"time"
)

var ErrNil = errors.New("查询结果为空")

var NoDelete int8 = 0 // 未删除

// Time is alias type for time.Time
type Time time.Time

const (
	timeFormat = rgconst.GoTimeFormat
	zone       = "Asia/Shanghai"
)

func NowTime() (t Time) {
	return Time(time.Now())
}

// UnmarshalJSON implements json unmarshal interface.
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// MarshalJSON implements json marshal interface.
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t Time) Int() int64 {
	return time.Time(t).Unix()
}

func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(zone)
	return time.Time(t).In(loc)
}

// Value ...
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time 注意是指针类型 method
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// SearchParam <LiJunDong : 2022-08-16 23:20:46> --- 查询入参结构
type SearchParam struct {
	This   *rgrequest.Client
	Query  string
	Args   []interface{}
	Order  string
	Fields []string
	Offset int
	Limit  int
}
