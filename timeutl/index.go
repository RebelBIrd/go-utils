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
	return strutl.ConnString(strconv.Itoa(y), "-", m.String(), "-", strconv.Itoa(d))
}
