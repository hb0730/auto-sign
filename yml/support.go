package config

type Support interface {
	//Do 核心执行
	Do(interface{})
	//DoVoid 默认由*AutoSign 调用，所有子类重写
	DoVoid()
	//Run 为*cron.Job，由*AutoSign重写
	Run()
	//GetConfig 获取其配置文件
	GetConfig(AutoSignConfig, int) interface{}
	//Supports 是否支持，返回具体的类型
	Supports(AutoSignConfig) Support
}
