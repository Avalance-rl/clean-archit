package pgx

import (
	"clean-artchit/internal/domain/entity"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// оперирует entity принимает и возвращает
type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (pr *productRepository) GetOne(ctx context.Context, id string) (entity.Product, error) {
	panic("implement me")
}

func (pr *productRepository) GetAll(ctx context.Context) ([]entity.Product, error) {
	panic("implement me")
}

func (pr *productRepository) Create(ctx context.Context, product *entity.Product) error {
	panic("implement me")
}

func (pr *productRepository) Update(ctx context.Context, product *entity.Product) error {
	panic("implement me")
}

func (pr *productRepository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
