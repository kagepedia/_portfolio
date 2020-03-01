package model

import "time"

type Task struct {
	ID        int64     `json:"id"`
	Task      string    `json:"task"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
