package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
	ErrMsg  string      `json:"error,omitempty"`
}

func respondWithErrorStatus(c echo.Context, status int, errMsg string) error {
	return c.JSON(status, &Response{Success: false, ErrMsg: errMsg})
}

func respondWithData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, &Response{Success: true, Data: data})
}
