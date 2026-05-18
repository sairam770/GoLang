package storage

import (
	"fmt"
	"sync"
)

// InMemoryStorage is a simple implementation of Storage using a map

type InMemoryStorage struct {
	users map[string]User
	mu    sync.RWMutex
}

// NewMemoryUserStorage initializes a new InMemoryStorage instance

func NewMemoryUserStorage() *InMemoryStorage {
	return &InMemoryStorage{
		users: make(map[string]User),
	}
}

// AddUser adds a new user to the storage

func (s *InMemoryStorage) AddUser(user User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.users[user.Id] = user
	return nil
}

//Get user fetches a user record or returns an error if the user does not exist

func (s *InMemoryStorage) GetUser(id string) (User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, exists := s.users[id]
	if !exists {
		return User{}, fmt.Errorf("user with id %s not found", id)
	}
	return user, nil
}
