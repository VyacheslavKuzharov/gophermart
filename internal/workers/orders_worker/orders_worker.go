package ordersworker

import (
	"context"
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/config"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/order"
	"github.com/VyacheslavKuzharov/gophermart/internal/workers/orders_worker/accrual"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"sync"
	"time"
)

type OrdersWorker struct {
	Queue      chan *Task
	Polling    time.Duration
	wg         *sync.WaitGroup
	cancelFunc context.CancelFunc
	repo       *order.Repository
	logger     *logger.Logger
	cfg        *config.Config
	accrual    *accrual.Accrual
}

func New(polling time.Duration, repo *order.Repository, logger *logger.Logger, cfg *config.Config) *OrdersWorker {
	return &OrdersWorker{
		Queue:   make(chan *Task),
		Polling: polling,
		repo:    repo,
		logger:  logger,
		wg:      new(sync.WaitGroup),
		cfg:     cfg,
		accrual: accrual.New(cfg.AccrualAddr, logger),
	}
}

func (ow *OrdersWorker) Start() {
	log := ow.logger.Logger
	pctx := context.Background()
	ctx, cancelFunc := context.WithCancel(pctx)
	ow.cancelFunc = cancelFunc

	log.Info().Msg("Run background workers...")

	ow.wg.Add(1)
	go ow.spawnDBPoller(ctx)

	for i := 0; i < 1; i++ {
		ow.wg.Add(1)
		go ow.spawnWorker(ctx)
	}

	//for i := 0; i < runtime.NumCPU(); i++ {
	//	ow.wg.Add(1)
	//	go ow.spawnWorker(ctx)
	//}
}

func (ow *OrdersWorker) Stop() {
	log := ow.logger.Logger
	log.Info().Msg("Stop workers")

	close(ow.Queue)
	ow.cancelFunc()
	ow.wg.Wait()
	log.Info().Msg("all background jobs are stop")
}

func (ow *OrdersWorker) spawnWorker(ctx context.Context) {
	defer ow.wg.Done()
	fmt.Println("worker awaiting")

	for task := range ow.Queue {
		select {
		case <-ctx.Done():
			return
		default:
			ow.perform(ctx, task)
		}
	}
}

func (ow *OrdersWorker) enqueueTask(task *Task) {
	ow.Queue <- task
}
