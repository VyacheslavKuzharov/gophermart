package workers

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

type worker struct {
	workchan    chan workType
	workerCount int
	buffer      int
	wg          *sync.WaitGroup
	cancelFunc  context.CancelFunc
}

type WorkerIface interface {
	Start(pctx context.Context)
	Stop()
	QueueTask(task string, workDuration time.Duration) error
}

func New(workerCount, buffer int) WorkerIface {
	w := worker{
		workchan:    make(chan workType, buffer),
		workerCount: workerCount,
		buffer:      buffer,
		wg:          new(sync.WaitGroup),
	}

	return &w
}

func (w *worker) Start(pctx context.Context) {
	ctx, cancelFunc := context.WithCancel(pctx)
	w.cancelFunc = cancelFunc

	for i := 0; i < w.workerCount; i++ {
		w.wg.Add(1)
		go w.spawnWorkers(ctx, i)
	}
}

func (w *worker) Stop() {
	log.Println("stop workers")
	close(w.workchan)
	w.cancelFunc()
	w.wg.Wait()
	log.Println("all workers exited!")
}

func (w *worker) QueueTask(task string, workDuration time.Duration) error {
	if len(w.workchan) >= w.buffer {
		return ErrWorkerBusy
	}

	w.workchan <- workType{TaskID: task, WorkDuration: workDuration}

	return nil
}

func (w *worker) spawnWorkers(ctx context.Context, num int) {
	defer w.wg.Done()
	log.Println("-----spawnWorkers---->", num)

	for work := range w.workchan {
		log.Println("-----range w.workchan---->")

		select {
		case <-ctx.Done():
			return
		default:
			log.Println("---range w.workchan--default---->")
			w.doWork(ctx, work.TaskID, work.WorkDuration)
		}
	}
}

func (w *worker) doWork(ctx context.Context, task string, workDuration time.Duration) {
	log.Printf("task: %s, do some work now...", task)
	sleepContext(ctx, workDuration)
	log.Printf("task: %s, work completed!", task)
}

func sleepContext(ctx context.Context, sleep time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(sleep):
	}
}

type workType struct {
	TaskID       string
	WorkDuration time.Duration
}

var (
	ErrWorkerBusy = errors.New("workers are busy, try again later")
)
