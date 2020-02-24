package application

import (
	"context"

	"server/domain/model"
	"server/domain/repository"
)

type TaskApp interface {
	GetAll(context.Context) ([]*model.Task, error)
}

type taskApp struct {
	taskRepository repository.TaskRepository
}

func NewTaskApp(tr repository.TaskRepository) TaskApp {
	return &taskApp{
		taskRepository: tr,
	}
}

// GetAll : Book データを全件取得するためのアプリケーション
//  -> 本システムではあまりアプリケーション層の恩恵を受けれないが、もう少し大きなシステムになってくると、
//    「ドメインモデルの調節者」としての役割が見えてくる
func (ta taskApp) GetAll(ctx context.Context) (tasks []*model.Task, err error) {
	// dbの呼び出し
	tasks, err = ta.taskRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
