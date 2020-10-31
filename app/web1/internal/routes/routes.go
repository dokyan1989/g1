package routes

import (
	"net/http"

	apiV1 "github.com/dokyan1989/g1/app/web1/internal/api/v1"
	"github.com/dokyan1989/g1/app/web1/internal/handler"
	"github.com/dokyan1989/g1/app/web1/internal/store"
	"github.com/dokyan1989/g1/lib/router"
)

// RegisterRoutes ...
func RegisterRoutes(store store.Querier) *router.Router {
	r := router.NewRouter()
	h := handler.New()
	api := apiV1.NewAPIHandler(store)

	// Product
	r.Route(http.MethodGet, "/api/v1/products/([0-9]+)", h.Build(api.GetProduct))
	r.Route(http.MethodGet, "/api/v1/products", h.Build(api.ListProducts))
	r.Route(http.MethodPost, "/api/v1/products", h.Build(api.CreateProduct))

	return r
}
