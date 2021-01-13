package support

//ISuperJob 定时任务
type ISuperJob interface {
	//ReadYaml 支持读取yaml
	ReadYaml
	//Run job Run
	Run()
}
