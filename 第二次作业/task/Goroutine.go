package task

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
var Mute1, Mute2, Mute3 sync.WaitGroup

func Gotest() {
	Mute1.Add(2)
	go func() {
		for i := 1; i < 11; i++ {
			if i%2 == 1 {
				fmt.Printf("%d", i)
			}
		}
		Mute1.Done()
	}()
	go func() {
		for j := 2; j < 11; j++ {
			if j%2 == 0 {
				fmt.Printf("%d", j)
			}
		}
		Mute1.Done()
	}()
	Mute1.Wait()
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

type Job struct {
	JobId     int
	JobExcute func()
}

type JobScheduler struct {
	Jobs []Job
}

type JobResult struct {
	JobId      int
	ExcuteTime time.Duration
}

func (js *JobScheduler) AddJob(job Job) {
	js.Jobs = append(js.Jobs, job)
}

func (js *JobScheduler) RunJob() []JobResult {
	var jobresult []JobResult
	Mute2.Add(len(js.Jobs))
	for _, job := range js.Jobs {
		go func(Job) {
			start := time.Now()
			job.JobExcute()
			Duration := time.Since(start)
			jobresult = append(jobresult, JobResult{JobId: job.JobId, ExcuteTime: Duration})
			Mute2.Done()
		}(job)

	}
	Mute2.Wait()
	return jobresult
}
