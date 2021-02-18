# auto-sign
go 实现签到

## geekhub 
代码实现 [geekhub](https://geekhub.com) 的签到
首先需要一个原始的`session_id`
```go
cookie := make(map[string]string, 1)
cookie["_session_id"] = ""
geekhub := Geekhub{Cookies: cookie}
geekhub.Do()
```

只需要Cookie中的`_session_id`

## appletuan 
代码实现[appletuan](https://appletuan.com)的签到
首先需要一个原始的`session_id`
```go
tuan := AppleTuan{Cookies: map[string]string{
"_session_id": "",
}}
tuan.Do()
```

只需要Cookie中的`_session_id`


## ld246
代码实现 [ld](https://ld246.com) 的签到
```go
ld := LD{Username: "", Password: ""}
ld.Do()
```
## V2ex
代码实现 [V2ex](https://V2ex.com) 的签到()
```go
params := make(map[string]string, 2)
params["PB3_SESSION"] = ""
params["A2"] = ""
params["V2EX_LANG"]=""
v2 := V2ex{Cookies: params}
v2.Do()
```

需要Cookie中的`v2`,`PB3_SESSION`

# 依赖
* [rod](https://github.com/go-rod/rod) 用于checkin
* [yaml](https://github.com/go-yaml/yaml) ~~解析yaml文件~~
* [uber/config](https://github.com/uber-go/config) 解析yaml
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
    V2EX_LANG:   
cron:
  geekhub: "5 0 * * *"
  appletuan: "10 0 * * *"
  ld246: "3 0 * * *"
  v2ex: "5 8 * * *"
mail:
  enabled: false
  host: "smtp.qq.com"
  protocol: "smtp"
  port: 465
  username: "xxxx@qq.com"
  password: "xxxx"
  fromName: "auto-sign"
  to: "xxxx@xx.com"
```