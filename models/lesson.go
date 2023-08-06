package models

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	Id        uuid.UUID
	Title     string
	CreatedAt time.Time
	Questions Questions
}

type Lessons []*Lesson
