package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jmhammock/wand-go/infrastructure"
	"github.com/jmhammock/wand-go/repositories"
	"golang.org/x/crypto/bcrypt"
)

type authResource struct {
	repo repositories.IUserRepo
}

func NewAuthResource(repo repositories.IUserRepo) *authResource {
	return &authResource{
		repo: repo,
	}
}

func (ar *authResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", ar.login)

	return r
}

func (ar *authResource) login(w http.ResponseWriter, r *http.Request) {
	data := &loginRequest{}
	err := render.Bind(r, data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := ar.repo.Get(data.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := infrastructure.GenJwt(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	render.Render(w, r, newLoginResponse(token))
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lr *loginRequest) Bind(r *http.Request) error {
	if lr.Email == "" || lr.Password == "" {
		return errors.New("missing email or password")
	}

	return nil
}

type loginResponse struct {
	Token string `json:"token"`
}

func newLoginResponse(token string) *loginResponse {
	resp := &loginResponse{
		Token: token,
	}

	return resp
}

func (lr *loginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
