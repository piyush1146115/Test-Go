package main

import (
	"errors"
	"log"
)

func main() {

	// Create new MemoryDB
	database := NewMemDB()

	// create the API and apply the memory database to it
	api := API{
		db: database,
	}
	// Add User using the API instead
	err := api.db.AddUser("findMe", "findMe@nowhere.hidden")
	if err != nil {
		log.Fatal(err)
	}
	// fetch user will also still work using the interface
	user, err := api.db.GetUser("findMe")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}

// UserDatabase is an interface that describes what's expected from structs that should be used as databases
type UserDatabase interface {
	UserRetriever
	AddUser(username, email string) error
}

// UserRetriver is a broken out interface that will be embedded by UserDatabase
type UserRetriever interface {
	GetUser(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

// API is a struct that will act as an API and connect to a database
type API struct {
	db UserDatabase
}

// User is a data representation of our users
type User struct {
	Name  string
	Email string
}

// MemoryDatabase is a user database that holds users in memory
type MemoryDatabase struct {
	Users map[string]*User
}

func NewMemDB() *MemoryDatabase {
	return &MemoryDatabase{
		Users: make(map[string]*User),
	}
}

// GetUser will print a user or an Error if not found
func (memDB *MemoryDatabase) GetUser(username string) (*User, error) {
	// Search for User
	if user, ok := memDB.Users[username]; ok {
		return user, nil
	}
	return nil, errors.New("User does not exist")
}

// GetUserByEmail will fetch users by email instead
func (memDB *MemoryDatabase) GetUserByEmail(email string) (*User, error) {
	for _, user := range memDB.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("User does not exist")
}

// AddUser will add a user to the database
func (memDB *MemoryDatabase) AddUser(username, email string) error {
	if _, ok := memDB.Users[username]; ok {
		return errors.New("User already exists")
	}
	memDB.Users[username] = &User{
		Name:  username,
		Email: email,
	}
	return nil
}
