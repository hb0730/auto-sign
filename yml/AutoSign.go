package config

//AutoSign 主要是做一个Abstract类，用于切面
type AutoSign struct {
	Support
	// Sub 用于*Support.Run()调用,防止丢失当前类
	//
	// 用法示例:
	//g := Geekhub{}
	// g.Cookies = map[string]string{"test": "测"}
	// g.Sub = g
	Sub     Support
	SubName string
}

func (support AutoSign) Run() {
	// 抓取异常发送email
	if support.Sub != nil {
		support.Sub.DoVoid()
	}
}
