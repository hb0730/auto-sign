# auto-sign
go 实现签到

## geekhub
代码实现 [geekhub](https://geekhub.com) 的签到
首先需要一个原始的`session_id`
```go
	geekhub := geekhub.Geekhub{SessionId: ""}
	geekhub.Do()
```
## ld246
代码实现 [ld](https://ld246.com) 的签到(自动签到)
```go
	ld := ld246.LD{Username: "", Password: ""}
	ld.Do()
```