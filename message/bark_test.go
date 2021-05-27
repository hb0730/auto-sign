package message

import "testing"

func TestBark_Send(t *testing.T) {
	m := GetSupport()
	if m != nil {
		msg := MessageBody{Title: "测试", Content: "测试错误"}
		m.Send(msg)
	}
}
