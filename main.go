package main

import (
	"fmt"
	"fuckandgo/controllers"
	"fuckandgo/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView *views.View
	contactView *views.View
	faqView *views.View
)


func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func page404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 NERD FUCKED UP DUDE</h1>")
}

func must(err error){
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")

	usersC := controllers.NewUsers()

	var h http.Handler = http.HandlerFunc(page404)
	r := mux.NewRouter()
	r.NotFoundHandler = h
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}
