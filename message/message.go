package message

import "github.com/hb0730/auto-sign/config"

var Messages = make(map[string]Message, 0)

// Message 消息发送
type Message interface {
	// Send 发送消息
	Send(MessageBody)
}

// GetSupport 获取支持
func GetSupport() Message {
	yaml := config.ReadYaml()
	key := yaml.String("message.type")
	return Messages[key]
}

// Enabled 是否启用
func Enabled() bool {
	yaml := config.ReadYaml()
	return yaml.Bool("message.enabled")
}

// Register 注册
func Register(name string, message Message) {
	Messages[name] = message
}

type MessageBody struct {
	Title   string
	Content string
}
