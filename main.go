package main

import (
	"auto-sign/util"
	config "auto-sign/yml"
	"github.com/robfig/cron/v3"
)

type Jobs struct {
	contextId cron.EntryID
	jobName   string
	cron      string
}

var jobs = make(map[string]Jobs)
var supportJob = map[string]interface{}{
	"geekhub": config.Geekhub{},
	"ld246":   config.Ld{},
	"v2ex":    config.V2ex{},
}
var support = []config.Support{config.Geekhub{}, config.Ld{}, config.V2ex{}}

func main() {
	c := cron.New()
	_, err := c.AddFunc("* * * * *", func() {
		autoSign, err := config.RedStruct()
		if err != nil {
			util.ErrorF("%v\n", err)
			c.Stop()
		}
		expressionMap := autoSign.Cron
		if len(expressionMap) == 0 {
			return
		}
		// 循环表达式
		for k, v := range expressionMap {
			job, ok := jobs[k]
			//新添加的
			if !ok {
				do(k, v, c)
				// 已存在 ,且cron已修改
			} else if ok && job.cron != v {
				job := jobs[k]
				c.Remove(job.contextId)
				do(k, v, c)
			}
		}
	})
	if err != nil {
		util.ErrorF("%v \n", err)
		return
	}
	// 其中任务
	c.Start()
	// 关闭任务
	defer c.Stop()
	select {}
}

func do(k string, v string, c *cron.Cron) {
	// 所支持的
	if supportJob, ok := supportJob[k]; ok {
		id, err := c.AddFunc(v, func() {
			for _, v := range support {
				err := v.Support(supportJob)
				if err == nil {
					v.Do()
				}
			}
		})
		if err == nil {
			jobs[k] = Jobs{contextId: id, jobName: k, cron: v}
		}
	}

}
