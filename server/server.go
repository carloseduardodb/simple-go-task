package server

import (
	"go_task/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/tasks/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateTask(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodPost)

	r.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetAllTasks(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet)

	r.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetOneTask(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet)

	r.HandleFunc("/tasks/edit/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			controllers.EditTask(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodPut)

	r.HandleFunc("/tasks/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			controllers.DeleteTask(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodDelete)

	return r
}
