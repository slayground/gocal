package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	// ErrNotFound -> no user in the db
	ErrNotFound = errors.New("models: resource not found")
)

// NewUserService -> create a new instance of an UserService
// with error and db handling
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	// dev mode only
	db.LogMode(true)
	// defer db.Close() -> dont use here
	return &UserService{
		db: db,
	}, nil
}

// UserService -> service struct
type UserService struct {
	db *gorm.DB
}

// ByID -> what might happen
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil,otherError (something else went wrong -> 500 error)
func (us *UserService) ByID(id uint) (*User, error) {
	var user User

	err := us.db.Where("id = ?", id).First(&user).Error

	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create -> create provided user
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// Close -> closes the server database connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// DestructiveReset -> drops the user table and rebuilds it
// dev only
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

// User -> user model to be stored in database
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
