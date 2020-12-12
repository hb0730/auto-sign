# auto-sign
go 实现签到

## geekhub 已验证
代码实现 [geekhub](https://geekhub.com) 的签到
首先需要一个原始的`session_id`
```go
cookie := make(map[string]string, 1)
cookie["_session_id"] = ""
geekhub := Geekhub{Cookies: cookie}
geekhub.Do()
```

只需要Cookie中的`_session_id`

## ld246
代码实现 [ld](https://ld246.com) 的签到(自动签到)
```go
ld := LD{Username: "", Password: ""}
ld.Do()
```
## V2ex
代码实现 [V2ex](https://V2ex.com) 的签到()
```go
params := make(map[string]string, 2)
v2 := V2ex{cookies: params}
v2.Do()
```

需要Cookie中的`v2`,`PB3_SESSION`

# 依赖
* [rod](https://github.com/go-rod/rod)

# cron 定时任务
可以使用 [gocron](https://github.com/go-co-op/gocron) 来做定时任务