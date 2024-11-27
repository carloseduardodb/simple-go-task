package services

import (
	"go_task/database"
	"go_task/dto"
	"go_task/models"
	"log"
)

func EditTask(id int, payload *dto.Task) (*models.Task, error) {
	var task models.Task

	err := database.Connection().Get(&task, "UPDATE tasks SET text = $1, completed = $2 WHERE id = $3 RETURNING id, text, completed, created_at", payload.Text, payload.Completed, id)

	if err != nil {
		log.Println("Failed to update task", err)
		return nil, err
	}

	return &task, nil
}
