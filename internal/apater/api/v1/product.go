package v1

import "C"
import (
	"clean-artchit/internal/apater/api/dto"
	"clean-artchit/internal/domain/entity"
	usecaseProduct "clean-artchit/internal/domain/usecase/product"
	"context"
	"encoding/json"
	"net/http"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, product usecaseProduct.CreateProductDTO) (entity.Product, error)
	GetProductByID(ctx context.Context, id string) (entity.Product, error)
}

type productHandler struct {
	productUsecase ProductUsecase
}

func NewHandler(productUsecase ProductUsecase) *productHandler {
	return &productHandler{productUsecase: productUsecase}
}

func (h *productHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /products", h.createProduct)
}

func (h *productHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product := usecaseProduct.CreateProductDTO{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		InStock:     req.InStock,
	}

	createdProduct, err := h.productUsecase.CreateProduct(r.Context(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}
