package server

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	apiq "github.com/stalin-777/test-apiq"
	"github.com/stalin-777/test-apiq/config"
	"github.com/stalin-777/test-apiq/http"
	"github.com/stalin-777/test-apiq/logger"
	"github.com/stalin-777/test-apiq/queue"
)

type Server struct {
	Router *echo.Echo

	TaskService *queue.TaskService
}

//Run - Launches server and worker pool
func Run(cfg *config.Config) {

	server, err := getRouter(cfg)
	if err != nil {
		logger.Fatalf("Failed to ger router, error:%s", err.Error())
	}

	tasksToWorkers := make(chan *apiq.Task, cfg.WorkersNum)

	for i := 0; i < cfg.WorkersNum; i++ {
		go worker(tasksToWorkers)
	}

	go server.serveWorker(tasksToWorkers)

	routerSocket := fmt.Sprintf("%s:%v", cfg.Web.Host, cfg.Web.Port)
	logger.Infof("Service is running on socket %v\n", routerSocket)
	if err := server.Router.Start(routerSocket); err != nil {
		logger.Fatalf("Failed to start service, error:%s", err.Error())
	}
}

func getRouter(cfg *config.Config) (*Server, error) {

	s := &Server{
		Router: echo.New(),
		TaskService: &queue.TaskService{
			Queue: queue.New(cfg.TTL),
		},
	}
	s.registerHandlers()

	s.Router.Use(middleware.Recover())

	return s, nil
}

func (s *Server) registerHandlers() {

	var h http.Handler

	s.Router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API status: online")
	})

	h.TaskService = s.TaskService

	s.Router.GET("/tasks", h.FindTasks)
	s.Router.POST("/tasks", h.CreateTask)
}

func worker(tasks <-chan *apiq.Task) {

	for {
		queue.DoTask(<-tasks)
	}

}

func (s *Server) serveWorker(tasks chan *apiq.Task) {

	for {

		if s.TaskService.Queue.Len() > 0 {
			tasks <- s.TaskService.Queue.Pop()
		} else {
			time.Sleep(time.Second)
		}
	}
}
