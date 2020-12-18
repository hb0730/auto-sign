package main

import (
	"auto-sign/util"
	config "auto-sign/yml"
	"github.com/robfig/cron/v3"
	"sync"
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
	util.Info("start ..............")
	var wg sync.WaitGroup
	wg.Add(1)
	c := cron.New()
	_, err := c.AddFunc("0 */12 * * *", func() {
		autoSign, err := config.RedStruct()
		util.InfoF("%v \n", autoSign)
		if err != nil {
			util.ErrorF("%v\n", err)
			c.Stop()
			wg.Done()
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
				do(k, v, c, autoSign)
				// 已存在 ,且cron已修改
			} else if ok && job.cron != v {
				job := jobs[k]
				c.Remove(job.contextId)
				do(k, v, c, autoSign)
			}
		}
	})
	if err != nil {
		util.ErrorF("%v \n", err)
		wg.Done()
	}
	// 其中任务
	c.Start()
	// 关闭任务
	defer c.Stop()
	wg.Wait()
}

// k 为yaml corn key
// v 为yaml corn value
// c 为*cron.Cron 定时任务
// autoSign 为 config.AutoSign 配置参数
func do(k string, v string, c *cron.Cron, autoSign config.AutoSign) {
	// 所支持的
	if supportJob, ok := supportJob[k]; ok {
		id, err := c.AddFunc(v, func() {
			for _, v := range support {
				err := v.Support(supportJob)
				if err == nil {
					v.Do(autoSign)
				}
			}
		})
		if err == nil {
			jobs[k] = Jobs{contextId: id, jobName: k, cron: v}
		}
	}

}
