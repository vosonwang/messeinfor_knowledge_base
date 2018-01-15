package model

import (
	"github.com/satori/go.uuid"
	"time"
)

type Base struct {
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
	DeletedAt *time.Time `json:"deleted"`
}
