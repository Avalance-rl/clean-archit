package product

import "context"

type Repository interface {
	GetOne(ctx context.Context, id string) (*Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
}
