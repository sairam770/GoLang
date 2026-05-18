package main

import (
	"fmt"

	"github.com/interface/storage"
)

type UserService struct {
	storage *storage.FileStorage
}

// func main() {
// 	inMemoryStorage := storage.NewMemoryUserStorage()
// 	service := UserService{storage: inMemoryStorage}

// 	// Test adding a user
// 	user := storage.User{Id: "usr_1", Name: "Alice", Age: 30}

// 	_ = service.storage.AddUser(user)

// 	// Test retrieving a user

// 	fetchedUser, err := service.storage.GetUser("usr_1")

// 	if err != nil {
// 		fmt.Println("Error fetching user:", err)
// 		return
// 	}
// 	fmt.Printf("Fetched User: %+v\n", fetchedUser)
// }

func main() {

	// 1. Initialize the new file backend

	fileBackend, err := storage.NewFileStorage("./data_store")
	if err != nil {
		panic("Failed to initialize file storage: " + err.Error())
	}

	// 2. Inject it into the service

	service := UserService{storage: fileBackend}

	//3. Perform actions

	newUser := storage.User{Id: "usr_2", Name: "Bob", Age: 25}

	err = service.storage.AddUser(newUser)
	if err != nil {
		fmt.Println("Error adding user:", err)
		return
	}
	fmt.Println("User added successfully.")

	// 4. Retrieve the data from disk

	fetched, err := fileBackend.GetUser("usr_22")
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return
	}
	fmt.Printf("Fetched User: %+v\n", fetched)
}
