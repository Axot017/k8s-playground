package handler

import (
	"net/http"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Register(r *http.ServeMux) {
	r.HandleFunc("/health", h.getHealth)
}

func (h *Health) getHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}
