package service

import (
	"clean-artchit/internal/domain/entity"
	"clean-artchit/internal/domain/usecase/product"
	"context"
)

// Сервисы оперируют доменной сущностью

type ProductRepository interface {
	GetOne(ctx context.Context, id string) (entity.Product, error)
	GetAll(ctx context.Context) ([]entity.Product, error)
	Create(ctx context.Context, product entity.Product) error
	Update(ctx context.Context, product entity.Product) error
	Delete(ctx context.Context, id string) error
}

type productService struct {
	repository ProductRepository
}

func NewService(repository ProductRepository) *productService {
	return &productService{repository: repository}
}

func (s *productService) GetProductByID(ctx context.Context, id string) (entity.Product, error) {
	return s.repository.GetOne(ctx, id)
}

func (s *productService) GetAllProducts(ctx context.Context, limit, offset int) ([]entity.Product, error) {
	return s.repository.GetAll(ctx)
}

func (s *productService) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	panic("implement me")
}

func (s *productService) UpdateProduct(ctx context.Context, product product.UpdateProductDTO) (*entity.Product, error) {
	panic("implement me")
}
