package application

import (
	"os"
	"testing"
)

func TestLd246_Start(t *testing.T) {
	ld := Ld246{Username: os.Getenv("username"), Password: os.Getenv("password")}
	ld.Start()
}
