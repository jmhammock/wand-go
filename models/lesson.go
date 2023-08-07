package models

import (
	"time"
)

type Lesson struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Questions Questions `json:"questions"`
}

type Lessons []*Lesson
