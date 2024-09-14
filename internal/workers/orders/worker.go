package orders

import "fmt"

type Worker struct {
	id    int
	queue *Scheduler
}

func NewWorker(id int, queue *Scheduler) *Worker {
	return &Worker{
		id:    id,
		queue: queue,
	}
}

func (w *Worker) Spawn() {
	for {
		//t := w.queue.PopWait()
		//
		//err := w.resizer.Resize(t.Filename)
		//if err != nil {
		//	fmt.Printf("error: %v\n", err)
		//	continue
		//}

		fmt.Printf("worker #%d\n", w.id)

		w.queue.PopWait()
	}
}
