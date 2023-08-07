package models

import (
	"time"
)

type Option struct {
	Id         int64     `json:"id"`
	OptionType string    `json:"option_type"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}

type Options []*Option
