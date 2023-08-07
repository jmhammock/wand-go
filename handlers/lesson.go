package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jmhammock/wand-go/models"
	"github.com/jmhammock/wand-go/repositories"
)

type adminLessonResource struct {
	repo repositories.ILessonRepo
}

func NewAdminLessonResource(repo repositories.ILessonRepo) *adminLessonResource {
	return &adminLessonResource{
		repo: repo,
	}
}

func (ar *adminLessonResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/lessons", func(r chi.Router) {
		r.Get("/", ar.list)
		r.Get("/{id}", ar.get)
	})

	return r
}

func (ar *adminLessonResource) list(w http.ResponseWriter, r *http.Request) {
	lessons, err := ar.repo.List()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	render.RenderList(w, r, newLessonListResponse(*lessons))
}

func (ar *adminLessonResource) get(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	if param == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lesson, err := ar.repo.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	render.Render(w, r, newLessonResponse(lesson))
}

type lessonResponse struct {
	*models.Lesson
}

func newLessonResponse(lesson *models.Lesson) *lessonResponse {
	return &lessonResponse{
		Lesson: lesson,
	}
}

func newLessonListResponse(lessons models.Lessons) []render.Renderer {
	list := []render.Renderer{}
	for _, lesson := range lessons {
		list = append(list, newLessonResponse(lesson))
	}
	return list
}

func (lr *lessonResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
