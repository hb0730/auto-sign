package support

import (
	"github.com/hb0730/auto-sign/config"
	"testing"
)

func TestFamijia_DoRun(t *testing.T) {
	config.LoadKoanf("")
	f := Famijia{}
	f.Name = "famijia"
	f.ISupport = f
	tests := []struct {
		name    string
		fields  Famijia
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", f, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.DoRun(); (err != nil) != tt.wantErr {
				t.Errorf("DoRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
