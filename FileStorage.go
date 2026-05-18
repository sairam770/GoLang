package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// FileStorage implements the Storage interface using a JSON file for persistence

type FileStorage struct {
	dir string
}

// NewFileStorage initializes a new FileStorage instance with the specified directory

func NewFileStorage(dir string) (*FileStorage, error) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}
	return &FileStorage{dir: dir}, nil
}

// buildPath returns the unique file path for a given user ID

func (s *FileStorage) buildPath(id string) string {
	return filepath.Join(s.dir, fmt.Sprintf("%s.json", id))
}

//  Save user data to a JSON file named after the user's ID

func (s *FileStorage) AddUser(user User) error {
	path := s.buildPath(user.Id)
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write user file: %w", err)
	}
	return nil
}

// reads the json file from disk and unmarshals it into a User struct

func (s *FileStorage) GetUser(id string) (User, error) {
	path := s.buildPath(id)
	data, errors := os.ReadFile(path)
	if errors != nil {
		return User{}, fmt.Errorf("failed to read user file: %w", errors)
	}

	var user User
	errors = json.Unmarshal(data, &user)
	if errors != nil {
		return User{}, fmt.Errorf("failed to unmarshal user: %w", errors)
	}

	return user, nil
}
