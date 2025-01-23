package db

import (
	"clean-artchit/internal/domain/product"
	"context"
)

type productRepository struct {
}

func (pr *productRepository) GetOne(ctx context.Context, id string) (*product.Product, error) {
	panic("implement me")
}

func (pr *productRepository) GetAll(ctx context.Context) ([]product.Product, error) {
	panic("implement me")
}

func (pr *productRepository) Create(ctx context.Context, product *product.Product) error {
	panic("implement me")
}

func (pr *productRepository) Update(ctx context.Context, product *product.Product) error {
	panic("implement me")
}

func (pr *productRepository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
