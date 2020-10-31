package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dokyan1989/g1/app/web1/internal/handler"
	"github.com/dokyan1989/g1/app/web1/internal/store"
	"github.com/dokyan1989/g1/lib/router"
)

// GetProduct ...
func (h *APIHandler) GetProduct(w http.ResponseWriter, r *http.Request) error {
	urlParams := router.URLParams(r)
	id, err := strconv.Atoi(urlParams[0])
	if err != nil {
		return err
	}

	product, err := h.store.GetProduct(r.Context(), uint(id))
	if err != nil {
		return handler.StatusError{Code: http.StatusNotFound, Err: err}
	}

	handler.WriteSuccess(w, product, http.StatusOK)
	return nil
}

// ListProducts ...
func (h *APIHandler) ListProducts(w http.ResponseWriter, r *http.Request) error {
	var listProductsParams store.ListProductsParams
	err := router.QueryParams(r, &listProductsParams)
	if err != nil {
		return err
	}
	
	product, err := h.store.ListProducts(r.Context(), listProductsParams)
	if err != nil {
		return handler.StatusError{Code: http.StatusNotFound, Err: err}
	}

	handler.WriteSuccess(w, product, http.StatusOK)
	return nil
}

// CreateProduct ...
func (h *APIHandler) CreateProduct(w http.ResponseWriter, r *http.Request) error {
	var createProductParams store.CreateProductParams
	err := router.BodyParams(r, &createProductParams)
	if err != nil {
		return err
	}

	err = h.store.CreateProduct(r.Context(), createProductParams)
	if err != nil {
		return err
	}

	fmt.Println(createProductParams)
	handler.WriteSuccess(w, apiResponse{Message: "The product was created successfully"}, http.StatusCreated)
	return nil
}
