package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Fucked up moron on the death side.</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in fuck, please fuck an email "+
		"to <a href=\"mailto:udon@example.com\">"+
		"udon@example.com</a>.")
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQU FUCKU FUCK YOU morons!</h1>"+
		"<h2>Are you sure to die? [Y]/y</h2>")
}

func page404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 NERD FUCKED UP DUDE</h1>")
}

func main() {
	var h http.Handler = http.HandlerFunc(page404)
	r := httprouter.New()
	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)
	r.NotFound = h

	http.ListenAndServe(":3000", r)
}
