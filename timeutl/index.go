package timeutl

import (
	"github.com/qinyuanmao/go-utils/strutl"
	"strconv"
	"time"
)

func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetDate() string {
	y, m, d := time.Now().Date()
	return strutl.ConnString(strconv.Itoa(y), "-", strconv.Itoa(int(m)), "-", strconv.Itoa(d))
}

type Time time.Time

const timeFormat = "2006-01-02 15:04:05"

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
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

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}
