package logutl

import (
	"runtime"
	"strconv"
	"time"

	"github.com/gookit/color"

	"github.com/qinyuanmao/go-utils/timeutl"
)

func Error(message ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Red.Println("error:"+timeutl.Time(time.Now()).String()+":"+f.Name()+":"+strconv.Itoa(line)+":", message)
	}
}

func Debug(message ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Green.Println("debug:"+timeutl.Time(time.Now()).String()+":"+f.Name()+":"+strconv.Itoa(line)+":", message)
	}
}

func Info(message ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.White.Println("info:"+timeutl.Time(time.Now()).String()+":"+f.Name()+":"+strconv.Itoa(line)+":", message)
	}
}

func Warning(message ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Yellow.Println("warning:"+timeutl.Time(time.Now()).String()+":"+f.Name()+":"+strconv.Itoa(line)+":", message)
	}
}
