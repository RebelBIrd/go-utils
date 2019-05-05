package logutl

import (
	"fmt"
	"runtime"
)

const (
	TextRed    = 31
	TextGreen  = 32
	TextYellow = 33
	TextWhite  = 37
)

func Red(str ...interface{}) interface{} {
	return textColor(TextRed, str)
}

func Green(str ...interface{}) interface{} {
	return textColor(TextGreen, str)
}

func Yellow(str ...interface{}) interface{} {
	return textColor(TextYellow, str)
}

func White(str ...interface{}) interface{} {
	return textColor(TextWhite, str)
}

func textColor(color int, str ...interface{}) interface{} {
	if IsWindows() {
		return str
	}
	switch color {
	case TextRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
	case TextGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
	case TextYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
	case TextWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
	default:
		return str
	}
}

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}
