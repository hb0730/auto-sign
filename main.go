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

var yamlConfig config.YamlConfig

func main() {
	util.Info("start ..............")
	var wg sync.WaitGroup
	wg.Add(1)
	c := cron.New()
	_, err := c.AddFunc("30 * * * *", func() {
		var err error
		yamlConfig, err = config.RedStruct()
		util.InfoF("%v \n", yamlConfig)
		if err != nil {
			util.ErrorF("%v\n", err)
			c.Stop()
			wg.Done()
		}
		expressionMap := yamlConfig.Cron
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
		wg.Done()
	}
	// 其中任务
	c.Run()
	// 关闭任务
	defer c.Stop()
	wg.Wait()
}

// k 为yaml corn key
// v 为yaml corn value
// c 为*cron.Cron 定时任务
func do(k string, v string, c *cron.Cron) {
	// 所支持的
	if supportJob, ok := config.SupportsMap[k]; ok {
		job := supportJob.(config.Support)
		job = job.Supports(yamlConfig)
		id, err := c.AddJob(v, job)
		if err == nil {
			jobs[k] = Jobs{contextId: id, jobName: k, cron: v}
		}
	}
}
