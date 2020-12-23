package config

type Support interface {
	//Do 核心执行
	Do(interface{})
	//DoVoid 默认由*AbstractSupport 调用，所有子类重写
	DoVoid()
	//Run 为*cron.Job，由*AutoSign重写
	Run()
	//Supports 是否支持，返回具体的类型
	Supports(YamlConfig) Support
}
