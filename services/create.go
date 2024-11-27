package services

import (
	"go_task/database"
	"go_task/dto"
	"go_task/models"
	"log"
)

func CreateTask(payload *dto.Task) (*models.Task, error) {
	var task models.Task

	err := database.Connection().Get(&task, "INSERT INTO tasks (text, completed, created_at) VALUES ($1, $2, NOW()) RETURNING id, text, completed, created_at", payload.Text, payload.Completed)

	if err != nil {
		log.Println("Failed to create task", err)
		return nil, err
	}

	return &task, nil
}
