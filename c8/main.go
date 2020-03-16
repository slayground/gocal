package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/phamstack/gocal/models"
)

// constants declaration
const (
	host = "localhost"
	port = 5432
	user = "huypham"
	// password = "pw"
	dbname = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	// use when we need to reset out database
	// us.DestructiveReset()

	// user := models.User{
	// 	Name:  "Dan James",
	// 	Email: "danjames@mu.co",
	// }

	// if err := us.Create(&user); err != nil {
	// 	panic(err)
	// }

	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
