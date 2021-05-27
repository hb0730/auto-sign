package support

import (
	"github.com/go-rod/rod"
	"github.com/hb0730/auto-sign/utils"
	"time"
)

//支持的类型
var Supports = make(map[string]AutoRun, 0)

// ISupport 实际签到
type ISupport interface {
	DoRun() error
}

//Support 支持的类型
// 承上启下作用
type Support struct {
	// AutoRun Cron调用
	AutoRun
	// ISupport 后置
	// 需要注册
	// hub := Geekhub{}
	//	hub.ISupport = hub
	ISupport
}

// Run Cron执行
func (s Support) Run() {
	utils.Info("cron 开始执行")
	err := rod.Try(func() {
		retry(s, 3)
	})
	if err != nil {
		utils.Error(err.Error())
	}
}

// Register 将支持的类型进行注册
func Register(name string, support AutoRun) {
	Supports[name] = support
}

// retry 尝试机制
func retry(a Support, num int) {
	err := rod.Try(func() {
		e := a.DoRun()
		if e != nil {
			panic(e)
		}
	})
	if err != nil && num > 0 {
		for {
			num--
			time.Sleep(time.Duration(3) * time.Second)
			retry(a, num)
		}
	} else if err != nil {
		panic(err)
	}
}
