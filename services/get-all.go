package services

import (
	"fmt"
	"go_task/database"
	"go_task/models"
)

func GetAllTasks(filters map[string]string) (*[]models.Task, error) {
	var tasks []models.Task

	if filters["completed"] != "true" && filters["completed"] != "false" && filters["completed"] != "" {
		return nil, fmt.Errorf("invalid completed value: %s", filters["completed"])
	}

	err := database.Connection().Select(&tasks, "SELECT id, text, completed, created_at FROM tasks WHERE completed = $1", filters["completed"] == "true")

	if err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []models.Task{}
	}

	return &tasks, nil
}
