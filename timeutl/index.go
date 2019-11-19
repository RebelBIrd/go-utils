package timeutl

import (
	"fmt"
	"github.com/qinyuanmao/go-utils/strutl"
	"strconv"
	"time"
)

type Time time.Time

const timeFormat = "2006/01/02 15:04:05"
const timeFormat2 = "2006-01-02 15:04:05"

func GetNowTime() string {
	return time.Now().Format(timeFormat)
}

func GetTimeString(t time.Time) string {
	return t.Format(timeFormat)
}

func GetTimeString2(t time.Time) string {
	return t.Format(timeFormat2)
}

func GetDate() string {
	y, m, d := time.Now().Date()
	return strutl.ConnString(strconv.Itoa(y), "-", strconv.Itoa(int(m)), "-", strconv.Itoa(d))
}

func GetOldDate(date int64) string {
	y, m, d := time.Unix(date, 0).Date()
	return strutl.ConnString(strconv.Itoa(y), "-", strconv.Itoa(int(m)), "-", strconv.Itoa(d))
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	if err != nil {
		fmt.Println(err)
	}
	*t = Time(now)
	return
}

func (t *Time) UnmarshalString(data string) (err error) {
	now, err := time.ParseInLocation(timeFormat, data, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func Now() Time {
	return Time(time.Now())
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}
