# auto-sign
go 实现签到

## geekhub 已验证
代码实现 [geekhub](https://geekhub.com) 的签到
首先需要一个原始的`session_id`
```go
geekhub := geekhub.Geekhub{SessionId: ""}
geekhub.Do()
```

只需要Cookie中的`_session_id`

## ld246 -未验证
代码实现 [ld](https://ld246.com) 的签到(自动签到)
```go
ld := ld246.LD{Username: "", Password: ""}
ld.Do()
```
## V2ex -未验证
代码实现 [V2ex](https://V2ex.com) 的签到()
```go
v := v2ex.V2ex{Cookie: ""}
v.Do()
```

需要Cookie中的`v2`,`PB3_SESSION`

# cron 定时任务
可以使用 [gocron](https://github.com/go-co-op/gocron) 来做定时任务