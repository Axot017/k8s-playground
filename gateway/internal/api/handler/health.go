package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Register(r chi.Router) {
	r.Get("/health", h.getHealth)
}

func (h *Health) getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
