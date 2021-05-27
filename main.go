package main

import (
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/support"
	"github.com/hb0730/auto-sign/utils"
	"github.com/robfig/cron/v3"
	"sync"
)

//Jobs 用户记录运行的job
type Jobs struct {
	contextId cron.EntryID
	jobName   string
	cron      string
}

var jobs = make(map[string]Jobs)

//ReadCron 读取cron表达式
func ReadCron() (Cron, error) {
	v := config.ReadYaml()
	r := v.GetStringMapString("cron")
	return Cron{Cron: r}, nil
}

//Cron Cron struct
type Cron struct {
	Cron map[string]string
}

func main() {
	utils.Info("main start ....")
	var wg sync.WaitGroup
	wg.Add(1)
	c := cron.New()
	//每30分钟读取配置文件
	_, err := c.AddFunc("* * * * *", func() {
		readCron, e := ReadCron()
		//如果读取异常，则关闭守护
		if e != nil {
			utils.ErrorF(e.Error())
			c.Stop()
			wg.Done()
			return
		}
		if len(readCron.Cron) == 0 {
			return
		}
		for k, v := range readCron.Cron {
			doJob(k, v, c)
		}
	})

	if err != nil {
		utils.ErrorF("%v\n", err)
		wg.Done()
	}
	// 其中任务
	c.Run()
	// 关闭任务
	defer c.Stop()

	wg.Wait()
}

func doJob(name string, value string, c *cron.Cron) {
	job, ok := jobs[name]
	//新添加的
	if !ok {
		do(name, value, c)
	} else if ok && job.cron != value {
		// 已存在 ,且cron已修改
		c.Remove(job.contextId)
		do(name, value, c)
	}
}

func do(k string, v string, c *cron.Cron) {
	run, ok := support.Supports[k]
	if !ok {
		return
	}
	id, err := c.AddJob(v, run)
	if err == nil {
		jobs[k] = Jobs{contextId: id, jobName: k, cron: v}
	}
}
