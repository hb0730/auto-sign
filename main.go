package main

import (
	"auto-sign/support/aggregate"
	cron2 "auto-sign/support/cron"
	"auto-sign/util"
	"github.com/robfig/cron/v3"
	"sync"
)

//Jobs 用户记录运行的job
type Jobs struct {
	contextId cron.EntryID
	jobName   string
	cron      string
}

//jobs 记录添加的Job key为支持的类型
var jobs = make(map[string]Jobs)

func main() {
	util.Info("start ......")
	var wg sync.WaitGroup
	wg.Add(1)
	c := cron.New()
	_, err := c.AddFunc("* * * * *", func() {
		support, err := cron2.Read()
		if err != nil {
			util.ErrorF("%v\n", err)
			c.Stop()
			wg.Done()
		}
		// 判断是否已有表达式
		if len(support.Cron) <= 0 {
			return
		}

		for k, v := range support.Cron {
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

// do 具体执行的操作
// supportName 为yaml corn key
// cornValue 为yaml corn value
// c 为*cron.Cron 定时任务
func do(supportName string, cornValue string, c *cron.Cron) {
	supportType := aggregate.GetSupports(supportName)
	if supportType == -1 {
		return
	}
	supportJob := aggregate.NewInstance(supportType)
	job, err := supportJob.Read()
	if err == nil {
		id, err := c.AddJob(cornValue, job)
		if err == nil {
			jobs[supportName] = Jobs{contextId: id, jobName: supportName, cron: cornValue}
		}
	}
}
