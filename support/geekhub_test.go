package support

import (
	"github.com/hb0730/auto-sign/config"
	"testing"
)

func TestGeekhub_DoRun(t *testing.T) {
	config.LoadKoanf("")
	g := Geekhub{}
	g.Name = "geekhub"
	g.ISupport = g
	tests := []struct {
		name    string
		fields  Geekhub
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", g, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.DoRun(); (err != nil) != tt.wantErr {
				t.Errorf("DoRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
