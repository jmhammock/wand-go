package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jmhammock/wand-go/handlers"
	"github.com/jmhammock/wand-go/repositories"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	lessonRepo := repositories.NewLessonRepo(db)
	userRepo := repositories.NewUserRepo(db)

	r.Mount("/admin", handlers.NewAdminLessonResource(lessonRepo).Routes())
	r.Mount("/auth", handlers.NewAuthResource(userRepo).Routes())

	log.Fatal(http.ListenAndServe(":8080", r))
}
