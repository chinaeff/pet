// Package pet предоставляет структуры данных для работы с питомцами.
package petRepo

import (
	"errors"
	"sync"
)

type PetRepo struct {
	Mu         sync.RWMutex
	DataID     map[int64]Pet
	DataStatus map[string]Pet
}

// Pet представляет собой информацию о питомце.
// swagger:model
type Pet struct {
	ID        int64     `json:"id"`
	Category  *Category `json:"category"`
	Name      string    `json:"name"`
	PhotoURLs []byte    `json:"photoURLs"`
	Tags      *Tag      `json:"tags"`
	Status    string    `json:"status"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewPetRepo() *PetRepo {
	return &PetRepo{DataID: make(map[int64]Pet)}
}

func (r *PetRepo) GetPetByID(ID int64) (*Pet, error) {
	r.Mu.RLock()
	defer r.Mu.Unlock()

	if pet, ok := r.DataID[ID]; ok {
		return &pet, nil
	}
	return nil, errors.New("No petRepo")
}
func (r *PetRepo) GetPetByStatus(status string) (*Pet, error) {
	r.Mu.RLock()
	defer r.Mu.Unlock()

	if pet, ok := r.DataStatus[status]; ok {
		return &pet, nil
	}
	return nil, errors.New("No petRepo")
}

func (r *PetRepo) DeletePetByID(ID int64) (*Pet, error) {
	r.Mu.RLock()
	defer r.Mu.Unlock()

	if _, ok := r.DataID[ID]; ok {
		delete(r.DataID, ID)
		return nil, errors.New("Pet deleted")
	}
	return nil, errors.New("Pet not found")

}

func (r *PetRepo) PutPetByID(ID int64, updatedPet Pet) error {
	r.Mu.RLock()
	defer r.Mu.Unlock()

	if _, ok := r.DataID[ID]; ok {
		r.DataID[ID] = updatedPet
		return nil
	}
	return errors.New("Pet not found")
}

func (r *PetRepo) PostPetByID(ID int64, newPet Pet) error {
	r.Mu.RLock()
	defer r.Mu.Unlock()

	if _, ok := r.DataID[ID]; !ok {
		r.DataID[ID] = newPet
		return nil
	}
	return errors.New("Failed to add new pet")
}
func (r *PetRepo) PostImageByID(ID int64, image []byte) error {
	r.Mu.RLock()
	defer r.Mu.Unlock()
	if pet, ok := r.DataID[ID]; ok {
		pet.PhotoURLs = image
		r.DataID[ID] = pet
		return nil
	}
	return errors.New("No petRepo")
}
func (r *PetRepo) PostPet(newPet Pet) error {
	r.Mu.RLock()
	defer r.Mu.Unlock()

	r.DataID[newPet.ID] = newPet

	return nil
}
