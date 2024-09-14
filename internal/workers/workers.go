package workers

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/config"
	"github.com/VyacheslavKuzharov/gophermart/internal/workers/orders"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"runtime"
)

func Run(cfg *config.Config, l *logger.Logger) {
	l.Logger.Info().Msg("Run background workers...")

	scheduler := orders.NewScheduler(5)
	scheduler.Start(context.Background())

	scheduler.Queue <- &orders.Task{Msg: "qweqweqweqeqwe"}

	for i := 0; i < runtime.NumCPU(); i++ {
		worker := orders.NewWorker(i, scheduler)
		go worker.Spawn()
	}
}
