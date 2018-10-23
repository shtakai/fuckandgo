package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "fuckandgo_dev"
)

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserId uint
	Amount int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	fmt.Println("connected to DB")

	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var user User
	db.Preload("Orders").First(&user)
	if db.Error != nil {
		panic(db.Error)
	}

	fmt.Println("email:", user.Email)
	fmt.Println("name:", user.Name)
	fmt.Println("orders:", user.Orders)

}

//func createOrder(db *gorm.DB, user User, amount int, desc string) {
//	db.Create(&Order{
//		UserId: user.ID,
//		Amount: amount,
//		Description: desc,
//	})
//	if db.Error != nil {
//		panic(db.Error)
//	}
//}
