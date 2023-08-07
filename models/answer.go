package models

import "time"

type Answer struct {
	QuestionId int64     `json:"question_id"`
	UserId     int64     `json:"user_id"`
	Option     Option    `json:"option"`
	CreatedAt  time.Time `json:"created_at"`
}
