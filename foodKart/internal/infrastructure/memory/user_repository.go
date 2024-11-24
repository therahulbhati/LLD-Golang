package memory

import (
	"errors"
	"sync"

	"foodkart/internal/domain/user"
)

type UserRepository struct {
	users map[string]*user.User
	mu    sync.RWMutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*user.User),
	}
}

func (r *UserRepository) Save(u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[u.ID] = u
	return nil
}

func (r *UserRepository) FindByID(id string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) FindByPhoneNumber(phoneNumber string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, user := range r.users {
		if user.PhoneNumber == phoneNumber {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
