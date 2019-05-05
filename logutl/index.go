package logutl

import "fmt"

func Error(message ...interface{}) {
	fmt.Println(Red("--------------Error Start--------------"))
	fmt.Println(Red(message))
	fmt.Println(Red("--------------Error End--------------"))
}

func Debug(message ...interface{}) {
	fmt.Println(Green("--------------Debug Start--------------"))
	fmt.Println(Green(message))
	fmt.Println(Green("--------------Debug End--------------"))
}

func Info(message ...interface{}) {
	fmt.Println(White("--------------Info Start--------------"))
	fmt.Println(White(message))
	fmt.Println(White("--------------Info End--------------"))
}

func Warning(message ...interface{}) {
	fmt.Println(Yellow("--------------Warning Start--------------"))
	fmt.Println(Yellow(message))
	fmt.Println(Yellow("--------------Warning End--------------"))
}
