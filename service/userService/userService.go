package userService

import (
	"pet/repository/userRepo"
)

type UserService struct {
	repo *userRepo.UserRepo
}

func NewUserService(repo *userRepo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Get(name string) (userRepo.User, error) {
	return s.repo.Get(name)
}

func (s *UserService) Put(user userRepo.User) error {
	return s.repo.Put(user)
}

func (s *UserService) PostNewUser(newUser userRepo.User) error {
	return s.repo.PostNewUser(newUser)
}

func (s *UserService) GetLogin(username, password string) (userRepo.User, error) {
	return s.repo.GetLogin(username, password)
}

func (s *UserService) GetLogout(username string) error {
	return s.repo.GetLogout(username)
}

func (s *UserService) DeleteUser(username string) error {
	return s.repo.DeleteUser(username)
}

func (s *UserService) PostNewArrayOfUsers(users []userRepo.User) error {
	return s.repo.PostNewArrayOfUsers(users)
}

func (s *UserService) PostNewListOfUser(users ...userRepo.User) error {
	return s.repo.PostNewListOfUser(users...)
}
