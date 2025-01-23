package product

import (
	"clean-artchit/internal/apater/api"
	"clean-artchit/internal/domain/product"
	"encoding/json"
	"net/http"
)

type handler struct {
	productService product.Service
}

func NewHandler(productService product.Service) api.Handler {
	return &handler{productService: productService}
}

func (h *handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /products", h.createProduct)
}

func (h *handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var req product.CreateProductDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.productService.CreateProduct(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}
