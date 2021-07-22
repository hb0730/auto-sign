package support

import (
	"github.com/hb0730/auto-sign/config"
	"testing"
)

func TestPagodaWxMini_DoRun(t *testing.T) {
	config.LoadKoanf("")
	pagoda := PagodaWxMini{}
	pagoda.Name = "pagodaWxMini"
	pagoda.ISupport = pagoda
	tests := []struct {
		name    string
		fields  PagodaWxMini
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", pagoda, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.DoRun(); (err != nil) != tt.wantErr {
				t.Errorf("DoRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
