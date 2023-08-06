package models

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id        uuid.UUID
	LessonId  uuid.UUID
	Text      string
	CreatedAt time.Time
	Options   Options
}

type Questions []*Question
