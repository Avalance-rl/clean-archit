package service

import "clean-artchit/internal/infrastructure/repository"

type Service struct {
	storeRepo   repository.StoreRepository
	productRepo repository.ProductRepository
}

func NewService(
	storeRepo repository.StoreRepository,
	productRepo repository.ProductRepository,
) *Service {
	return &Service{
		storeRepo:   storeRepo,
		productRepo: productRepo,
	}
}
