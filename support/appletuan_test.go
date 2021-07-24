package support

import (
	"github.com/hb0730/auto-sign/config"
	"testing"
)

func TestAppleTuan_DoRun(t *testing.T) {
	config.LoadKoanf("")
	tuan := AppleTuan{}
	tuan.Name = "appletuan"
	tuan.ISupport = tuan
	tests := []struct {
		name    string
		fields  AppleTuan
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", tuan, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.DoRun(); (err != nil) != tt.wantErr {
				t.Errorf("DoRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
