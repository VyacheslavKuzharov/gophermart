package scheduler

import (
	"fmt"
	"time"
)

type Job interface {
	Execute()
}

type PrintJob struct {
	Message string
}

func (p PrintJob) Execute() {
	fmt.Println(p.Message)
}

type JobScheduler struct {
	JobQueue chan Job
	Interval time.Duration
}

func NewJobScheduler(interval time.Duration) *JobScheduler {
	return &JobScheduler{
		JobQueue: make(chan Job),
		Interval: interval,
	}
}

func (s *JobScheduler) Start() {
	go func() {
		ticker := time.NewTicker(s.Interval)

		for {
			select {
			case job := <-s.JobQueue:
				fmt.Println("----job := <-s.JobQueue--->")
				job.Execute()
			case <-ticker.C:
				fmt.Println("----<-ticker.C--->")
				//for job := range s.JobQueue {
				//	fmt.Println("----job--->")
				//	job.Execute()
				//}
			}
		}
	}()
}

func (s *JobScheduler) ScheduleOnce(duration time.Duration, job Job) {
	go func() {
		time.Sleep(duration)
		s.JobQueue <- job
	}()
}
