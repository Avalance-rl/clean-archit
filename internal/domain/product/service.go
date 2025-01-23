package product

import "context"

type Service interface {
	GetProductByID(ctx context.Context, id string) (*Product, error)
	GetAllProducts(ctx context.Context, limit, offset int) ([]Product, error)
	CreateProduct(ctx context.Context, product CreateProductDTO) (*Product, error)
	UpdateProduct(ctx context.Context, product UpdateProductDTO) (*Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetProductByID(ctx context.Context, id string) (*Product, error) {
	return s.repository.GetOne(ctx, id)
}

func (s *service) GetAllProducts(ctx context.Context, limit, offset int) ([]Product, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) CreateProduct(ctx context.Context, product CreateProductDTO) (*Product, error) {
	panic("implement me")
}

func (s *service) UpdateProduct(ctx context.Context, product UpdateProductDTO) (*Product, error) {
	panic("implement me")
}
