package tasker

import (
	"context"
	"sync"
	"time"

	apiq "github.com/stalin-777/test-apiq"
	"github.com/stalin-777/test-apiq/internal/logger"
	"github.com/stalin-777/test-apiq/internal/queue"
	"github.com/stalin-777/test-apiq/internal/router"
)

type Tasker struct {
	TasksToWorkers chan *apiq.Task
	Ctx            context.Context
	CancelFunc     context.CancelFunc
	WG             *sync.WaitGroup
}

func New(workersNum int) *Tasker {

	tasker := &Tasker{}

	tasker.TasksToWorkers = make(chan *apiq.Task, workersNum)
	tasker.Ctx, tasker.CancelFunc = context.WithCancel(context.Background())
	tasker.WG = &sync.WaitGroup{}

	return tasker
}

func (s *Tasker) Run(server *router.Router, workersNum int) {

	go s.serveWorker(server)

	s.WG.Add(workersNum)
	for i := 0; i < workersNum; i++ {
		go s.worker()
	}
}

func (s *Tasker) serveWorker(server *router.Router) {

	for {

		if server.TaskService.Queue.Len() > 0 {

			// task := server.TaskService.Queue.Pop()
			// logger.Infof("Task id %s is queued", task.ID)
			s.TasksToWorkers <- server.TaskService.Queue.Pop()
		} else {

			select {
			case <-s.Ctx.Done():
				close(s.TasksToWorkers)

				return
			default:
				time.Sleep(time.Second)
			}
		}
	}
}

func (s *Tasker) worker() {

	defer s.WG.Done()

	for task := range s.TasksToWorkers {

		logger.Infof("Task id %v started", task.ID)
		queue.DoTask(task)
		logger.Infof("Task id %v completed", task.ID)
	}
}
