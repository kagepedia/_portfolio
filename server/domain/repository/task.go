package repository

import (
	"context"

	"server/domain/model"
)

type TaskRepository interface {
	GetAll(context.Context) ([]*model.Task, error)
}
