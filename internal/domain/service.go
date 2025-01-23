package domain

import "clean-artchit/internal/domain/model"

type ServiceStore interface {
	GetStore() (model.Store, error) // фильтры подумать
	GetByIDStore(id string) (model.Store, error)
	GetStores() ([]model.Product, error)
	CreateStore(product *model.Store) error
	UpdateStore(product *model.Store) error
	DeleteStore(product *model.Store) error
}

type ServiceProduct interface {
	GetProduct() (model.Product, error) // фильтры подумать
	GetByIDProduct(id string) (model.Product, error)
	GetProducts() ([]model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(product *model.Product) error
}
