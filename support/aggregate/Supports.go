package aggregate

import (
	"auto-sign/support"
	"auto-sign/support/appletuan"
	"auto-sign/support/geekhub"
	"auto-sign/support/ld246"
	"auto-sign/support/v2ex"
)

//const SupportStr;
type SupportType int

const (
	GEEKHUB SupportType = iota
	APPLETUAN
	LD246
	V2EX
)

var SupportTypes = [...]string{"geekhub", "appletuan", "ld246", "v2ex"}

//GetSupports 获取支持的类型
func GetSupports(types string) SupportType {
	for i, v := range SupportTypes {
		if v == types {
			return SupportType(i)
		}
	}
	return -1
}

var supportMaps = map[SupportType]support.ISuperJob{
	GEEKHUB:   geekhub.Geekhub{},
	APPLETUAN: appletuan.AppleTuan{},
	LD246:     ld246.Ld246{},
	V2EX:      v2ex.V2ex{},
}

//NewInstance 获取对应类型
func NewInstance(support SupportType) support.ISuperJob {
	return supportMaps[support]
}
