package config

type Support interface {
	Do(interface{})
	Run()
	GetConfig(AutoSignConfig, int) interface{}
	Support(AutoSignConfig) Support
}
