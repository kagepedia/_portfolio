package db

import (
	"database/sql"
	"log"
	"server/domain/model"

	_ "github.com/go-sql-driver/mysql"
)

type TaskRepository struct {
	DB *sql.DB
}

func (tr *TaskRepository) All() []*model.Task {
	var taskList []*model.Task
	rows, err := tr.DB.Query("SELECT * FROM t_task")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var task model.Task

		err := rows.Scan(&task)

		if err != nil {
			log.Fatal(err)
		}

		taskList = append(taskList, &task)
	}

	return taskList
}

/*
func (tr *TaskRepository) Save(task *model.Task) error {
	insert := "INSERT INTO t_task(task) VALUES(?)"

	if task.CreatedAt == "" {
		task.CreatedAt = "2015-05-05 07:53:54"
	}

	fmt.Println(task)

	_, err := tr.Conn.Exec(insert, task.Status, task.Task)

	return err
}
*/
