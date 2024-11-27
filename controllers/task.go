package controllers

import (
	"encoding/json"
	"go_task/dto"
	"go_task/handlers"
	"go_task/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task dto.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		handlers.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdTask, err := services.CreateTask(&task)
	if err != nil {
		handlers.SendError(w, http.StatusInternalServerError, "Failed to create task")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdTask); err != nil {
		log.Println("Failed to encode task", err)
		handlers.SendError(w, http.StatusInternalServerError, err.Error())
	}
}

func GetOneTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		handlers.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	task, err := services.GetOneTask(idInt)
	if err != nil {
		handlers.SendError(w, http.StatusInternalServerError, "Failed to get task")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Println("Failed to encode task", err)
		handlers.SendError(w, http.StatusInternalServerError, err.Error())
	}
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	params := map[string]string{"completed": r.URL.Query().Get("completed")}

	tasks, err := services.GetAllTasks(params)
	if err != nil {
		handlers.SendError(w, http.StatusInternalServerError, "Failed to get tasks")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println("Failed to encode tasks", err)
		handlers.SendError(w, http.StatusInternalServerError, err.Error())
	}
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		handlers.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	payload := dto.Task{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		handlers.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	task, err := services.EditTask(idInt, &payload)
	if err != nil {
		handlers.SendError(w, http.StatusInternalServerError, "Failed to edit task")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Println("Failed to encode task", err)
		handlers.SendError(w, http.StatusInternalServerError, err.Error())
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		handlers.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	services.DeleteTask(idInt)
}
