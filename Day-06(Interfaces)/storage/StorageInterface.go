package storage

import "errors"

// User defines the core data model

type User struct {
	Id   string
	Name string
	Age  int
}

// sentinal errors for clean error handling

var ErrUserNotFound = errors.New("user not found")

// Storage defines the interface for user data management

type Storage interface {
	AddUser(user User) error
	GetUser(id string) (User, error)
}
