package util

import (
	"fmt"
	"os"
)

// 来自https://github.com/wonderivan/logger
type brush func(string) string

func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

//鉴于终端的通常使用习惯，一般白色和黑色字体是不可行的,所以30,37不可用，
var colors = []brush{
	newBrush("1;31"), // Error              红色
	newBrush("1;33"), // Warn               黄色
	newBrush("1;36"), // Informational      天蓝色
	newBrush("1;32"), // Debug              绿色
	newBrush("1;32"), // Trace              绿色
}

func Error(msg string) {
	msg = colors[0](msg)
	out(msg)
}
func ErrorF(format string, v interface{}) {
	Error(fmt.Sprintf(format, v))
}

func Warn(msg string) {
	msg = colors[1](msg)
	out(msg)
}
func WarnF(format string, v interface{}) {
	Warn(fmt.Sprintf(format, v))
}

func Info(msg string) {
	msg = colors[2](msg)
	out(msg)
}

func InfoF(format string, v interface{}) {
	Info(fmt.Sprintf(format, v))
}

func out(msg string) {
	os.Stdout.Write(append([]byte(msg), '\n'))
}
