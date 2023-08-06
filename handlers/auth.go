package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmhammock/wand-go/repositories"
)

type AuthResource struct {
	repo repositories.IUserRepo
}

func (ar *AuthResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", ar.login)
	r.Post("/register", ar.register)

	return r
}

func (ar *AuthResource) login(w http.ResponseWriter, r *http.Request) {

}

func (ar *AuthResource) register(w http.ResponseWriter, r *http.Request) {

}
