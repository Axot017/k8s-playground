package handler

import "net/http"

type Handler interface {
	Register(*http.ServeMux)
}
