package application

import (
	"server/domain/model"
	"server/domain/repository"
)

type TaskApplication struct {
	Repo repository.TaskRepository
}

func (ta *TaskApplication) All() []*model.Task {
	return ta.Repo.All()
}

func (ta *TaskApplication) Save(t *model.Task) error {
	return ta.Repo.Save(t)
}
