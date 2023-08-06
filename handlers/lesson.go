package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jmhammock/wand-go/repositories"
)

type AdminResource struct {
	repo repositories.ILessonRepo
}

func (ar *AdminResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", ar.list)
	r.Get("/{id}", ar.get)

	return r
}

func (ar *AdminResource) list(w http.ResponseWriter, r *http.Request) {
	lessons, err := ar.repo.List()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(lessons)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(b)
}

func (ar *AdminResource) get(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id")
	if param == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	lesson, err := ar.repo.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(lesson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Write(b)
}
