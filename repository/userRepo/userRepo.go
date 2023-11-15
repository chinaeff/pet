package userRepo

import (
	"errors"
	"sync"
)

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userStatus"`
}

type UserRepo struct {
	mu   sync.RWMutex
	data map[string]User
}

func NewUserRepository() *UserRepo {
	return &UserRepo{
		data: make(map[string]User),
	}
}

func (r *UserRepo) Get(name string) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, ok := r.data[name]; ok {
		return user, nil
	}
	return User{}, errors.New("User not found")
}

func (r *UserRepo) Put(user User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[user.Username]; ok {
		r.data[user.Username] = user
		return nil
	}
	return errors.New("User not found")
}

func (r *UserRepo) PostNewUser(newUser User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[newUser.Username]; !ok {
		r.data[newUser.Username] = newUser
		return nil
	}
	return errors.New("User already exists")
}

func (r *UserRepo) GetLogin(username, password string) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.data[username]
	if !ok || user.Password != password {
		return User{}, errors.New("Invalid username or password")
	}

	return user, nil
}

func (r *UserRepo) GetLogout(username string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	//TODO: как показать выход из системы?

	return nil
}

func (r *UserRepo) DeleteUser(username string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[username]; ok {
		delete(r.data, username)
		return nil
	}
	return errors.New("User not found")
}

func (r *UserRepo) PostNewArrayOfUsers(users []User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range users {
		if _, ok := r.data[user.Username]; !ok {
			r.data[user.Username] = user
		}
	}

	return nil
}

func (r *UserRepo) PostNewListOfUser(users ...User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range users {
		if _, ok := r.data[user.Username]; !ok {
			r.data[user.Username] = user
		}
	}

	return nil
}
