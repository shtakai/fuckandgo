package main

import (
	"fmt"
	"fuckandgo/models"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "fuckandgo_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	us, err := models.NewUsersService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	user := models.User{
		Name: "世志凡太",
		Age: 98,
		Email: "fuck-seshi-bxxta@test.com",
	}

	if err := us.Create(&user); err != nil {
		panic(err)
	}

	user2 := models.User{
		Name: "世志凡太2",
		Age: 4,
		Email: "kil-seshi-bxxta@test.com",
	}

	if err := us.Create(&user2); err != nil {
		panic(err)
	}

	foundUser, err := us.ById(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)

	fmt.Println("==========")

	foundUser, err = us.ByEmail("fuck-seshi-bxxta@test.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)

	user.Name = "沖雅也"
	user.Email = "hikage-oyaji@nehande.matteru"

	err = us.Update(&user)
	if err != nil {
		panic(err)
	}

	foundUser, err = us.ByEmail("hikage-oyaji@nehande.matteru")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)

	fmt.Println("==========")

	foundUser, err = us.ByAge(user.Age)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)

	fmt.Println("==========")

	users := us.InAgeRange(1, 100)
    fmt.Println("users:", users)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(foundUser)

	fmt.Println("==========")

	err = us.Delete(foundUser.ID)
	if err != nil {
		panic(err)
	}

	_, err = us.ById(foundUser.ID)
	if err != models.ErrNotFound {
		panic("user is not deleted")
	}
	fmt.Println("User was fucked up RIPed")
}

