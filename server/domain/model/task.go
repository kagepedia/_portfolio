package model

import "time"

type Task struct {
	Id         int64
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
}
