package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	apiq "github.com/stalin-777/test-apiq"
	"github.com/stalin-777/test-apiq/internal/logger"
)

type Handler struct {
	TaskService apiq.TaskService
}

var (
	StatusOK         = http.StatusOK
	StatusBadRequest = http.StatusBadRequest
)

//Tasks - get list Tasks
func (h *Handler) FindTasks(c echo.Context) error {

	Tasks, err := h.TaskService.FindTasks()
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, StatusBadRequest, err.Error())
	}

	logger.Info("Successful attempt to get a list of Tasks")
	return respondWithData(c, Tasks)
}

//CreateTask - create a new row in DB
func (h *Handler) CreateTask(c echo.Context) error {

	Task := &apiq.Task{}

	err := c.Bind(Task)
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, StatusBadRequest, err.Error())
	}

	err = h.TaskService.CreateTask(Task)
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, StatusBadRequest, err.Error())
	}

	logger.Infof("Successful attempt to create a Task. ID:%v", Task.ID)
	return respondWithData(c, Task)
}
