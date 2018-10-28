package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
	Age uint
}

type UserService struct {
	db *gorm.DB
}

var (
	ErrNotFound = errors.New("models:resource not found")
	ErrInvalidID = errors.New("models:Id provided was invalid")
)

func NewUsersService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

func (us *UserService) Close() error {
	return us.db.Close()
}

func (us *UserService) DestructiveReset() error {
	if err := us.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return us.AutoMigrate()
}

func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) ById(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
        return nil, err
	}
	return &user, nil
}

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
    err := first(db, &user)
    return &user, err
}

func (us *UserService) ByAge(age uint) (*User, error) {
	var user User
	db := us.db.Where("age = ?", age)
	err := first(db, &user)
	return &user, err
}

func (us *UserService) InAgeRange(from, to uint) []User {
	var users []User
	db := us.db.Where("age >= ? and age <= ?", from, to)
	find(db, &users)
	return users
}

func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

func (us *UserService) Delete(id uint) error {
	if id == 0 {
        return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

func find(db *gorm.DB, dst interface{}) {
	db.Find(dst)
}
