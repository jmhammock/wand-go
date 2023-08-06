package models

import (
	"time"

	"github.com/google/uuid"
)

type OptionType int64

const (
	Distractor OptionType = iota
	Correct
)

func (o OptionType) String() string {
	switch o {
	case Distractor:
		return "distractor"
	case Correct:
		return "correct"
	}
	return "unknown"
}

type Option struct {
	Id         uuid.UUID
	OptionType OptionType
	Text       string
	CreatedAt  time.Time
}

type Options []*Option
