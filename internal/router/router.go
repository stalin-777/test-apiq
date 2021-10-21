package router

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stalin-777/test-apiq/internal/config"
	"github.com/stalin-777/test-apiq/internal/http"
	"github.com/stalin-777/test-apiq/internal/logger"
	"github.com/stalin-777/test-apiq/internal/queue"
)

type Router struct {
	router *echo.Echo

	TaskService *queue.TaskService
}

//Run - Launches router and worker pool
func (s *Router) Run(cfg *config.Config) {

	routerSocket := fmt.Sprintf("%s:%v", cfg.Web.Host, cfg.Web.Port)
	logger.Infof("Service is running on socket %v\n", routerSocket)
	if err := s.router.Start(routerSocket); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Failed to start service, error:%s", err.Error())
	}
}

func (s *Router) Shutdown() error {
	return s.router.Shutdown(context.Background())
}

func New(cfg *config.Config) (*Router, error) {

	s := &Router{
		router: echo.New(),
		TaskService: &queue.TaskService{
			Queue: queue.New(cfg.TTL),
		},
	}
	s.registerHandlers()

	s.router.Use(middleware.Recover())

	return s, nil
}

func (s *Router) registerHandlers() {

	var h http.Handler

	s.router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API status: online")
	})

	h.TaskService = s.TaskService

	s.router.GET("/tasks", h.FindTasks)
	s.router.POST("/tasks", h.CreateTask)
}
