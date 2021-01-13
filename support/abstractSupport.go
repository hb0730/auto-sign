package support

import (
	"auto-sign/util"
	"fmt"
)

//AbstractSupport 主要是做一个Abstract类，用于切面
type AbstractSupport struct {
	ISuperJob
	//Sub 做支持的类型
	Sub Support
	// 类型name
	SubName string
}

func (support AbstractSupport) Run() {
	if support.Sub == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			util.ErrorF("run error %v \n", r)
			m, err := Read()
			if err != nil {
				panic(r)
			}
			if m.Enabled {
				var c string
				subject := fmt.Sprintf("%s 执行出错", support.SubName)
				if t, ok := r.(error); ok {
					c = t.Error()
				}
				content := fmt.Sprintf("auto-sign在执行cron时出错，具体详情\n: %v\n", c)
				_ = m.Send(subject, content)
				panic(r)
			}
		}
	}()
	support.Sub.DoSupport()
}
