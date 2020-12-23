package appletuan

import (
	"testing"
)

// Cookies is null
func TestAppleTuan_Do(t *testing.T) {
	tuan := AppleTuan{Cookies: nil}
	tuan.Do()
}

// success
func TestAppleTuan_Do2(t *testing.T) {
	tuan := AppleTuan{Cookies: map[string]string{
		"_session_id": "",
	}}
	tuan.Do()
}
