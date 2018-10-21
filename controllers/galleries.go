package controllers

import (
	"fuckandgo/views"
	"net/http"
)

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/new"),
	}
}

type Galleries struct {
	NewView *views.View
}

// GET /new
func (u *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

