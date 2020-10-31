package v1

import (
	"github.com/dokyan1989/g1/app/web1/internal/store"
)

type apiResponse struct {
	Message string `json:"message,omitempty"`
}

// APIHandler ...
type APIHandler struct {
	store store.Querier
}

// NewAPIHandler ...
func NewAPIHandler(store store.Querier) *APIHandler {
	return &APIHandler{
		store: store,
	}
}
