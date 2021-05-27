package utils

import (
	"testing"
)

func TestError(t *testing.T) {
	Error("错误")
}

func TestErrorF(t *testing.T) {
	ErrorF("%s", "错误")
}

func TestInfo(t *testing.T) {
	Info("详情")
}

func TestInfoF(t *testing.T) {
	InfoF("%s", "详情")
}

func TestWarn(t *testing.T) {
	Warn("警告")
}

func TestWarnF(t *testing.T) {
	WarnF("%s", "警告")
}
