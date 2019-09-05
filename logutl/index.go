package logutl

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/qinyuanmao/go-utils/timeutl"
	"runtime"
	"strconv"
	"time"
)

func Error(message ...interface{}) {
	fmt.Println()
	color.Red.Println("------------------------------------------------ERROR-----------------------------------------------")
	color.Red.Println("Time: " + timeutl.Time(time.Now()).String())
	for i := 1; i <= 5; i ++ {
		pc, _, line, ok := runtime.Caller(i)
		f := runtime.FuncForPC(pc)
		if ok && f.Name() != "runtime.goexit" {
			color.Red.Println("Caller Skip" + strconv.Itoa(i) + ": " + f.Name() + ":" + strconv.Itoa(line) + "")
		} else {
			break
		}
	}
	color.Red.Println("----------------------------------------------------------------------------------------------------")
	color.Red.Println(message)
	color.Red.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println()
}

func Debug(message ...interface{}) {
	fmt.Println()
	color.Green.Println("------------------------------------------------DEBUG-----------------------------------------------")
	color.Green.Println("Time: " + timeutl.Time(time.Now()).String())
	for i := 1; i <= 5; i ++ {
		pc, _, line, ok := runtime.Caller(i)
		f := runtime.FuncForPC(pc)
		if ok && f.Name() != "runtime.goexit" {
			color.Green.Println("Caller Skip" + strconv.Itoa(i) + ": " + f.Name() + ":" + strconv.Itoa(line) + "")
		} else {
			break
		}
	}
	color.Green.Println("----------------------------------------------------------------------------------------------------")
	color.Green.Println(message)
	color.Green.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println()
}

func Info(message ...interface{}) {
	fmt.Println()
	color.White.Println("------------------------------------------------INFO------------------------------------------------")
	color.White.Println("Time: " + timeutl.Time(time.Now()).String())
	for i := 1; i <= 5; i ++ {
		pc, _, line, ok := runtime.Caller(i)
		f := runtime.FuncForPC(pc)
		if ok && f.Name() != "runtime.goexit" {
			color.White.Println("Caller Skip" + strconv.Itoa(i) + ": " + f.Name() + ":" + strconv.Itoa(line) + "")
		} else {
			break
		}
	}
	color.White.Println("----------------------------------------------------------------------------------------------------")
	color.White.Println(message)
	color.White.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println()
}

func Warning(message ...interface{}) {
	fmt.Println()
	color.Yellow.Println("------------------------------------------------WARN------------------------------------------------")
	color.Yellow.Println("Time: " + timeutl.Time(time.Now()).String())
	for i := 1; i <= 5; i ++ {
		pc, _, line, ok := runtime.Caller(i)
		f := runtime.FuncForPC(pc)
		if ok && f.Name() != "runtime.goexit" {
			color.Yellow.Println("Caller Skip" + strconv.Itoa(i) + ": " + f.Name() + ":" + strconv.Itoa(line) + "")
		} else {
			break
		}
	}
	color.Yellow.Println("----------------------------------------------------------------------------------------------------")
	color.Yellow.Println(message)
	color.Yellow.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println()
}
