# auto-sign
go 实现签到

## geekhub 

代码实现 [geekhub](https://geekhub.com) 的签到
首先需要一个原始的`_session_id`

```go
cookie := make(map[string]string, 1)
cookie["_session_id"] = ""
geekhub := Geekhub{Cookies: cookie}
geekhub.Start()
```


## appletuan 

代码实现[appletuan](https://appletuan.com)的签到
首先需要一个原始的`_session_id`

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
* [yaml/v1](github.com/spf13/viper)  用于读取yaml配置
* [cron](https://github.com/robfig/cron) 定时任务
* [mail](https://github.com/xhit/go-simple-mail) email发送

# Docker 
[hb0730/auto-sign](https://hub.docker.com/r/hb0730/auto-sign)

具体请看 `docker-compsoe`

# demo 
```yaml
cron:
  geekhub: "0 7 * * *"
  appletuan: "40 7 * * *"
  ld246: "5 0 * * *"
  v2ex: "0 8 * * *"
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
    A2:
    PB3_SESSION:
message:
  type: bark
  bark:
    url:
    key:
  mail:
    host:
    protocol:
    port: 465
    username:
    password:
    from_name:
    to:
```

## **注意**

[yaml/v1](github.com/spf13/viper) 在读取配置时大小写不区分 [issues/1014](https://github.com/spf13/viper/issues/1014)