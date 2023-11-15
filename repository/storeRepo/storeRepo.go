package storeRepo

import (
	"errors"
	"sync"
	"time"
)

type Store struct {
	ID       int64     `json:"id"`
	PetID    int64     `json:"petID"`
	Quantity int       `json:"quantity"`
	ShipDate time.Time `json:"shipDate"`
	Status   string    `json:"status"`
	Complete bool      `json:"complete"`
}

type StoreRepo struct {
	mu    sync.RWMutex
	data  map[int64]Store
	order map[int64]bool
}

func (r *StoreRepo) Inventory(status string) ([]Store, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []Store
	for _, store := range r.data {
		if store.Status == status {
			result = append(result, store)
		}
	}
	return result, nil
}

func (r *StoreRepo) GetOrder(ID int64) (*Store, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if store, ok := r.data[ID]; ok {
		return &store, nil
	}
	return nil, errors.New("Store not found")
}

func (r *StoreRepo) DeleteOrder(ID int64) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if _, ok := r.data[ID]; ok {
		delete(r.data, ID)
		return nil
	}
	return errors.New("Order not found")
}

func (r *StoreRepo) Order(PetID int64, Quantity int, ShipDate time.Time, Status string, Complete bool) (int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	orderID := generateUniqueID(r.order)

	newOrder := Store{PetID: PetID, Quantity: Quantity, ShipDate: ShipDate, Status: Status, Complete: Complete}

	r.data[orderID] = newOrder
	r.order[orderID] = true
	return orderID, nil
}

func generateUniqueID(orders map[int64]bool) int64 {
	var orderID int64
	for {
		orderID = time.Now().UnixNano()
		if _, exists := orders[orderID]; !exists {
			break
		}
	}

	return orderID
}
