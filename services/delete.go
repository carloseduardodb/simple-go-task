package services

import (
	"go_task/database"
	"log"
)

func DeleteTask(id int) {
	_, err := database.Connection().Exec("DELETE FROM tasks WHERE id = $1 RETURNING id, text, completed, created_at", id)

	if err != nil {
		log.Println("Failed to delete task", err)
	}
}
