package models

import (
	"time"
)

type Question struct {
	Id        int64     `json:"id"`
	LessonId  int64     `json:"lesson_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Options   Options   `json:"options"`
}

type Questions []*Question
