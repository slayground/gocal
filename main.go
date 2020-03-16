package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// constants declaration
const (
	host = "localhost"
	port = 5432
	user = "huypham"
	// password = "pw"
	dbname = "lenslocked_dev"
)

// User -> user model
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
	Color string
	// 1 -> many, 1 user can have many orders
	Orders []Order
}

// Order -> orders model
type Order struct {
	gorm.Model
	// no negative id -> unsigned integer -> twice as many available
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	// simply verify postgres and psqlInfo valid
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// create user table if not yet existed
	// db.DropTableIfExists(&User{})
	db.LogMode(true)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Order{})

	var u User

	// fmt.Println(db.First(&u)) -> even printing causes the database to change

	// Preload orders then fetch user, which contains Orders[]
	if err := db.Preload("Orders").Last(&u).Error; err != nil {
		panic(err)
	}

	// preload order data
	fmt.Println(u)
	fmt.Println(u.Orders) // -> [] if not db.Preload("Orders")

	// seeding orders
	// createOrder(db, u, 585, "Tesla Model 1")
	// createOrder(db, u, 396, "Tesla Model 3")
	// createOrder(db, u, 343, "Tesla Model 5")
	// createOrder(db, u, 919, "Tesla Model 7")

	// error handling
	// if err := db.Where("email = ?", "huynet@github.io").First(&u).Error; err != nil {
	// 	switch err {
	// 	case gorm.ErrRecordNotFound:
	// 		fmt.Println("No user found!")
	// 	default:
	// 		panic(err)
	// 	}
	// }
	// fmt.Println(u)
}

// create order and error handling
func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error
	if err != nil {
		panic(err)
	}
}
