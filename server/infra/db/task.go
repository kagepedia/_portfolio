package db

import (
	"context"
	"time"

	"server/domain/model"
	"server/domain/repository"
)

type taskData struct{}

func NewTaskData() repository.TaskRepository {
	return &taskData{}
}

// GetAll : DB から Task データを全件取得（TaskRepository インターフェースの GetAll() を実装したもの）
//  -> 本来は DB からデータを取得するが、簡略化のために省略（モックデータを返却）
func (td taskData) GetAll(context.Context) ([]*model.Task, error) {
	task1 := model.Task{}
	task1.Id = 1
	task1.Name = "taro"
	task1.CreateTime = time.Now().Add(-24 * time.Hour)
	task1.UpdateTime = time.Now().Add(-24 * time.Hour)

	task2 := model.Task{}
	task2.Id = 2
	task2.Name = "hanako"
	task2.CreateTime = time.Now().Add(-24 * 7 * time.Hour)
	task2.UpdateTime = time.Now().Add(-24 * 7 * time.Hour)

	return []*model.Task{&task1, &task2}, nil
}
