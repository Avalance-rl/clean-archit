package product

import (
	"clean-artchit/internal/domain/entity"
	"context"
	"go.uber.org/zap"
)

// бизнес сценарии системы
// тоже entity возвращают

type Service interface {
	GetProductByID(ctx context.Context, id string) (entity.Product, error)
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
}

// StoreService тут могут быть другие интерфейсы, которые используются
type StoreService interface {
	//GetByID(ctx context.Context, id string) entity.Author
}

type productUsecase struct {
	productService Service
	//authorService AuthorService
	//genreService  GenreService
	log *zap.Logger
}

func NewProductUsecase(productService Service, log *zap.Logger) *productUsecase {
	return &productUsecase{
		productService: productService,
		log:            log,
	}
}

func (uc *productUsecase) GetProductByID(ctx context.Context, id string) (entity.Product, error) {
	product, err := uc.productService.GetProductByID(ctx, id)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (uc *productUsecase) CreateProduct(ctx context.Context, product CreateProductDTO) (entity.Product, error) {
	productEntity := entity.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		InStock:     product.InStock,
	}
	return uc.productService.CreateProduct(ctx, productEntity)
}
