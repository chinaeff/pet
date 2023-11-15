package storeService

import (
	"pet/repository/storeRepo"
)

type StoreService struct {
	repo *storeRepo.StoreRepo
}

func (s *StoreService) Inventory(status string) ([]storeRepo.Store, error) {
	return s.repo.Inventory(status)
}
func (s *StoreService) GetOrder(ID int64) (*storeRepo.Store, error) {
	return s.repo.GetOrder(ID)
}
func (s *StoreService) DeleteOrder(ID int64) error {
	return s.repo.DeleteOrder(ID)
}

func (s *StoreService) Order(newStore storeRepo.Store) (*storeRepo.Store, error) {
	return &newStore, nil
}
