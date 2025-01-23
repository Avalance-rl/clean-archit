package service

import "clean-artchit/internal/domain/model"

func (s *Service) GetStore() (model.Store, error) {
	panic("implement me")
}
func (s *Service) GetByIDStore(id string) (model.Store, error) {
	panic("implement me")
}

func (s *Service) GetStores() ([]model.Store, error) {
	panic("implement me")
}

func (s *Service) CreateStore(product *model.Store) error {
	panic("implement me")
}

func (s *Service) UpdateStore(product *model.Store) error {
	panic("implement me")
}

func (s *Service) DeleteStore(product *model.Store) error {
	panic("implement me")
}
