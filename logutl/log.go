package logutl

import (
	"github.com/qinyuanmao/go-utils/fileutl"
	"github.com/qinyuanmao/go-utils/timeutl"
	"strings"
	"time"
)
var fileManager = fileutl.Manager{
	Path: "log/",
}

func init() {
	fileManager.Name = time.Now().Format("2006-1-2_15-04-05") + ".log"
	_ = fileManager.Create()
}

func WriteLog(msgType, msg string) {
	nowTime := timeutl.GetNowTime()
	content := strings.Join([]string{"======Log Start======\n", nowTime + "\n","======Message type======\n", msgType + "\n", "======Content Message======\n", msg + "\n", "======Log End======\n"}, "")
	fileManager.Write(content)
}