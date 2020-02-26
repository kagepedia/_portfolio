package model

import (
	"time"
)

type User struct {
	ID        uint32
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
