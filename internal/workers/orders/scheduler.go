package orders

import (
	"context"
	"fmt"
	"time"
)

type Task struct {
	Msg string
}

type Scheduler struct {
	Queue    chan *Task
	Interval time.Duration
}

func NewScheduler(interval time.Duration) *Scheduler {
	return &Scheduler{
		Queue:    make(chan *Task),
		Interval: interval,
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(s.Interval)

		for {
			select {
			case <-ctx.Done():
				return
			case job := <-s.Queue:
				fmt.Println("----<-s.Queue:--->", job)
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

func (s *Scheduler) PopWait() *Task {
	// получаем задачу
	return <-s.Queue
}
