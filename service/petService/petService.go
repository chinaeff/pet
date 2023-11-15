package petService

import (
	"pet/repository"
	"pet/repository/petRepo"
)

type PetService struct {
	Repo *petRepo.PetRepo
}

func NewPetService(repo *petRepo.PetRepo) *PetService {
	return &PetService{repo}
}
func (s *PetService) GetPetByID(ID int64) (*petRepo.Pet, error) {
	return s.Repo.GetPetByID(ID)
}
func (s *PetService) GetPetByStatus(status string) (*petRepo.Pet, error) {
	return s.Repo.GetPetByStatus(status)
}
func (s *PetService) DeletePetByID(ID int64) (*petRepo.Pet, error) {
	return s.Repo.DeletePetByID(ID)
}

func (s *PetService) PutPetByID(ID int64, updatedPet petRepo.Pet) error {
	existingPet, err := s.Repo.GetPetByID(ID)
	if err != nil {
		return err
	}

	if updatedPet.Category != nil {
		existingPet.Category = updatedPet.Category
	}
	if updatedPet.Name != "" {
		existingPet.Name = updatedPet.Name
	}
	if updatedPet.PhotoURLs != nil {
		existingPet.PhotoURLs = updatedPet.PhotoURLs
	}
	if updatedPet.Tags != nil {
		existingPet.Tags = updatedPet.Tags
	}
	if updatedPet.Status != "" {
		existingPet.Status = updatedPet.Status
	}
	return s.Repo.PutPetByID(ID, *existingPet)
}

func (s *PetService) PostPetByID(ID int64, formData repository.FormData) error {
	existingPet, err := s.Repo.GetPetByID(ID)
	if err != nil {
		return err
	}
	existingPet.Name = formData.Name
	existingPet.Status = formData.Status

	return s.Repo.PostPetByID(ID, *existingPet)
}
func (s *PetService) PostImageByID(ID int64, image []byte) (*petRepo.Pet, error) {
	existingPet, err := s.Repo.GetPetByID(ID)
	if err != nil {
		return nil, err
	}
	existingPet.PhotoURLs = image

	return &petRepo.Pet{ID: ID, PhotoURLs: image}, nil
}
func (s *PetService) PostPet(newPet petRepo.Pet) (*petRepo.Pet, error) {
	return &newPet, nil
}
