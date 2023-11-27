package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Register(r *echo.Group) {
	r.GET("/health", h.getHealth)
}

func (h *Health) getHealth(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
