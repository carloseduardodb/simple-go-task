package services

import (
	"go_task/database"
	"go_task/models"
)

func GetOneTask(id int) (*models.Task, error) {
	var task models.Task

	err := database.Connection().Get(&task, "SELECT id, text, completed, created_at FROM tasks WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return &task, nil
}
