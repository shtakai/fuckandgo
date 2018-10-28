package controllers

import (
	"fmt"
	"fuckandgo/models"
	"fuckandgo/views"
	"net/http"
)

func NewUsers(us *models.UserService) *Users {
	return &Users{
        NewView: views.NewView("bootstrap", "users/new"),
        us: us,
	}
}

type Users struct {
	NewView *views.View
	us *models.UserService
}

type SignupForm struct {
	Name string `schema: "name"`
	Age uint `shema:"age"`
	Email string `schema:"email"`
	Password string `schema:"password"`
}

// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	printForm(form)
	user := models.User{
		Name: form.Name,
		Age: form.Age,
        Email: form.Email,
	}
	if err := u.us.Create(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	fmt.Fprintln(w, "User is", user)
}

func printForm(form SignupForm) {
	fmt.Println("Name is", form.Name)
	fmt.Println("Age is", form.Age)
	fmt.Println("Email is", form.Email)
	fmt.Println("Password is", form.Password)
}