package support

import "testing"

func TestAppleTuan_Run(t *testing.T) {
	tuan := AppleTuan{}
	tuan.ISupport = tuan
	tuan.Name = "苹果团"
	tuan.Run()
}
