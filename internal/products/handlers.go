package products

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. CALL the serice → return products
	// 2. Return JSON in an HTTP response

	products := []string{"Hello", "World"}

	json.NewEncoder(w).Encode(products)
}
