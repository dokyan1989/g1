package handler

import (
	"net/http"
)

// AppMiddleware ...
type AppMiddleware func(AppHandleFunc) AppHandleFunc

// AppHandleFunc ...
type AppHandleFunc func(w http.ResponseWriter, r *http.Request) error

// AppHandler ...
type AppHandler struct {
	h           AppHandleFunc
	middlewares []AppMiddleware
}

// New ...
func New(middlewares ...AppMiddleware) *AppHandler {
	return &AppHandler{middlewares: append(([]AppMiddleware)(nil), middlewares...)}
}

// Build ...
func (h *AppHandler) Build(handleFunc AppHandleFunc) *AppHandler {
	return &AppHandler{h: handleFunc, middlewares: h.middlewares}
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.h == nil {
		return
	}
	for i := range h.middlewares {
		h.h = h.middlewares[len(h.middlewares)-1-i](h.h)
	}

	err := h.h(w, r)
	if err != nil {
		WriteError(w, err)
	}
}
