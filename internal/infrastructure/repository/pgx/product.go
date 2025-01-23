package pgx

import "clean-artchit/internal/domain/model"

func (p *Pgx) FindProduct() (model.Product, error) {
	panic("implement me")
}

func (p *Pgx) FindByIDProduct(id string) (model.Product, error) {
	panic("implement me")
}

func (p *Pgx) FindAllProduct() ([]model.Product, error) {
	panic("implement me")
}

func (p *Pgx) CreateProduct(product *model.Product) error {
	panic("implement me")
}

func (p *Pgx) UpdateProduct(product *model.Product) error {
	panic("implement me")
}

func (p *Pgx) DeleteProduct(product *model.Product) error {
	panic("implement me")
}
