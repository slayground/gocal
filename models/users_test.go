package models

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func testingUserService() (*UserService, error) {
	const (
		host   = "localhost"
		port   = 5432
		user   = "huypham"
		dbname = "lenslocked_test"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	us, err := NewUserService(psqlInfo)
	if err != nil {
		return nil, err
	}
	// testing -> disable log mode
	us.db.LogMode(false)
	// clear the users table
	// good tests should have no specific order
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T) {
	tus, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}
	user := User{
		Name:  "Dan James",
		Email: "danjames@mu.co",
	}

	err = tus.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID == 0 {
		t.Errorf("want id > 0, got %d", user.ID)
	}
	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
		t.Errorf("expected created recent")
	}
	if time.Since(user.UpdatedAt) > time.Duration(5*time.Second) {
		t.Errorf("expected updated recent")
	}
}
