package main

import (
	"fmt"

	"github.com/test/init_project/task"
)

func main() {
	// var p *int = new(int)
	// *p = 10
	// task.Add10(p)
	// fmt.Println(*p)

	// task.Gotest()

	var js *task.JobScheduler
	js = &task.JobScheduler{Jobs: []task.Job{}}
	var job1 task.Job = task.Job{
		JobId: 1,
		JobExcute: func() {
			println("job1")
		},
	}
	var job2 task.Job = task.Job{
		JobId: 2,
		JobExcute: func() {
			println("job2")
		},
	}
	var job3 task.Job = task.Job{
		JobId: 3,
		JobExcute: func() {
			println("job3")
		},
	}
	js.AddJob(job1)
	js.AddJob(job2)
	js.AddJob(job3)
	fmt.Println(js.RunJob())
}
