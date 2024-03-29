# auto-sign

go 实现签到

## command

`./auto-sign -SERVER_ADDRESS="" -SERVER_CRON="" -SERVER_CONFIG=“”`

## geekhub

代码实现 [geekhub](https://geekhub.com) 的签到 首先需要一个原始的`_session_id`

```go
cookie := make(map[string]string, 1)
cookie["_session_id"] = ""
geekhub := Geekhub{Cookies: cookie}
geekhub.Start()
```

## appletuan

代码实现[appletuan](https://appletuan.com)的签到 首先需要一个原始的`_session_id`

```go
tuan := AppleTuan{Cookies: map[string]string{
"_session_id": "",
}}
tuan.Start()
```

## ld246

代码实现 [ld](https://ld246.com) 的签到

```go
ld := LD{Username: "", Password: ""}
ld.Start()
```

## V2ex

代码实现 [V2ex](https://V2ex.com) 的签到

```go
params := make(map[string]string, 2)
params["PB3_SESSION"] = ""
params["A2"] = ""
v2 := V2ex{Cookies: params}
v2.Start()
```

需要Cookie中的`v2`,`PB3_SESSION`

## ChinaG

代码实现 [ChinaG](https://cc.ax/) 自动签到

```go
g := ChinaG{
"username",
"password",
}
_ = g.Start()
```

需要 username/password 即可

# Message 消息推送

消息推送支持 [mail](https://github.com/xhit/go-simple-mail) 与 [bark](https://github.com/Finb/Bark) ios 推送

## mail 推送

需要内容

```
host: 地址
protocol: 协议
port: 465 端口
username: 用户名
password: 密码
from_name: 类似昵称
to: 发送地址
```

## bark

需要内容

```
url: 地址
key: 密钥
```

# 依赖

* [rod](https://github.com/go-rod/rod) 用于checkin
* [yaml/v1](https://github.com/spf13/viper)  用于写yaml
* [koanf](https://github.com/knadh/koanf)  用于读取yaml(解决大小写敏感)
* [cron](https://github.com/robfig/cron) 定时任务
* [mail](https://github.com/xhit/go-simple-mail) email发送

# Docker

[hb0730/auto-sign](https://hub.docker.com/r/hb0730/auto-sign)

具体请看 `docker-compsoe`

# demo

```yaml
geekhub:
  cookies:
    _session_id:
appletuan:
  cookies:
    _session_id:
ld246:
  user:
    username:
    password:
v2ex:
  cookies:
    PB3_SESSION:
    A2:
famijia:
  headers:
    token:
    blackBox:
    deviceId:
chinaG:
  user:
    username:
    password:
cron:
  geekhub: "0 7 * * *"
  appletuan: "40 7 * * *"
  ld246: "5 0 * * *"
  v2ex: "0 8 * * *"
  famijia: "10 8 * * *"
  chinag: "20 8 * * *"
message:
  enabled: true
  type: "bark"
  bark:
    url:
    key:
  mail:
    host: "smtp.qq.com"
    protocol: "smtp"
    port: 465
    username:
    password:
    from_name: "auto-sign"
    to:
```

## Docker

```yaml
version: '3.1'
services:
  auto-sign:
    restart: always
    image: hb0730/auto-sign:latest
    environment:
      - TZ=Asia/Shanghai
    container_name: auto-sign
    volumes:
      - ./config:/app/config/
```

```tree
.
├── config
│   └── application.yml
├── docker-compose.yml

```

## **注意**

<del>[yaml/v1](github.com/spf13/viper) 在读取配置时大小写不区分 [issues/1014](https://github.com/spf13/viper/issues/1014) </del>