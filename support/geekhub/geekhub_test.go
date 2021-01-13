package geekhub

import (
	"testing"
)

func TestGeekhub_Read(t *testing.T) {
	g := Geekhub{}
	result, _ := g.Read()
	result.Run()
}
