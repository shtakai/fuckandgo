package main

import (
	"fmt"
	"fuckandgo/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func page404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 NERD FUCKED UP DUDE</h1>")
}

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	var h http.Handler = http.HandlerFunc(page404)
	r := mux.NewRouter()
	r.NotFoundHandler = h
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}
