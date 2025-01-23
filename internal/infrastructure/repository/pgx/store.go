package pgx

import "clean-artchit/internal/domain/model"

func (p *Pgx) FindStore() (model.Store, error) {
	panic("implement me")
}

func (p *Pgx) FindByIDStore(id string) (model.Store, error) {
	panic("implement me")
}

func (p *Pgx) FindAllStore() ([]model.Store, error) {
	panic("implement me")
}

func (p *Pgx) CreateStore(product *model.Store) error {
	panic("implement me")
}

func (p *Pgx) UpdateStore(product *model.Store) error {
	panic("implement me")
}

func (p *Pgx) DeleteStore(product *model.Store) error {
	panic("implement me")
}
