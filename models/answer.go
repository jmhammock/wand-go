package models

import "github.com/google/uuid"

type Answer struct {
	QuestionId uuid.UUID
	UserId     uuid.UUID
	Option     Option
}
