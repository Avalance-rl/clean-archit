package pgx

import (
	"context"

	"clean-artchit/internal/domain/entity"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// оперирует entity принимает и возвращает
type productRepository struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewProductRepository(db *pgxpool.Pool, log *zap.Logger) *productRepository {
	return &productRepository{
		db:  db,
		log: log,
	}
}

func (pr *productRepository) GetOne(ctx context.Context, id string) (entity.Product, error) {
	panic("implement me")
}

func (pr *productRepository) GetAll(ctx context.Context) ([]entity.Product, error) {
	panic("implement me")
}

func (pr *productRepository) Create(ctx context.Context, product entity.Product) error {
	panic("implement me")
}

func (pr *productRepository) Update(ctx context.Context, product entity.Product) error {
	panic("implement me")
}

func (pr *productRepository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
