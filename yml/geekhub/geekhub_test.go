package geekhub

import (
	"testing"
)

func TestGeekhub_DoVoid(t *testing.T) {
	g := Geekhub{}
	g.Cookies = map[string]string{"test": "测"}
	g.Sub = g
	g.Run()
}
