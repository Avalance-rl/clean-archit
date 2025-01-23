package repository

import "clean-artchit/internal/domain/model"

type StoreRepository interface {
	FindStore() (model.Store, error) // фильтры подумать
	FindByIDStore(id string) (model.Store, error)
	FindAllStore() ([]model.Store, error)
	CreateStore(product *model.Store) error
	UpdateStore(product *model.Store) error
	DeleteStore(product *model.Store) error
}

type ProductRepository interface {
	FindProduct() (model.Product, error) // фильтры подумать
	FindByIDProduct(id string) (model.Product, error)
	FindAllProduct() ([]model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(product *model.Product) error
}
