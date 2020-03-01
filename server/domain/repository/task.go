package repository

import "server/domain/model"

type TaskRepository interface {
	All() []*model.Task
	Save(*model.Task) error
}
