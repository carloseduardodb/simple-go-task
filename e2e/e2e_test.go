package e2e

import (
	"bytes"
	"encoding/json"
	"go_task/database"
	"go_task/dto"
	"go_task/server"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func startGoDotEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func resetDatabase() {
	database.Connection().Exec("DELETE FROM tasks")
	database.Connection().Exec("ALTER SEQUENCE tasks_id_seq RESTART WITH 1")
}

func startTestServer() *httptest.Server {
	server.StartServer()
	return httptest.NewServer(server.StartServer())
}

func TestMain(m *testing.M) {
	startGoDotEnv()
	resetDatabase()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateTask(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	task := dto.Task{
		Text:      "Test Task",
		Completed: false,
	}

	taskJSON, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}

	resp, err := http.Post(server.URL+"/tasks/create", "application/json", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status 201 Created, got %d", resp.StatusCode)
	}
}

func TestGetAllTasks(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestGetOneTask(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks/1")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestEditTask(t *testing.T) {

	server := startTestServer()
	defer server.Close()

	task := dto.Task{
		Text:      "Updated Task",
		Completed: true,
	}

	taskJSON, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, server.URL+"/tasks/edit/1", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestDeleteTask(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	req, err := http.NewRequest(http.MethodDelete, server.URL+"/tasks/delete/1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestCreateTask_InvalidJSON(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	invalidJSON := `{"text": "Test Task", "completed": "not a boolean"}`

	resp, err := http.Post(server.URL+"/tasks/create", "application/json", strings.NewReader(invalidJSON))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var errorResponse map[string]string
	err = json.NewDecoder(resp.Body).Decode(&errorResponse)
	assert.NoError(t, err)
	assert.Contains(t, errorResponse["message"], "Invalid request body")
}

func TestGetOneTask_InvalidID(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks/abc")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var errorResponse map[string]string
	err = json.NewDecoder(resp.Body).Decode(&errorResponse)
	assert.NoError(t, err)
	assert.Contains(t, errorResponse["message"], "Invalid request body")
}

func TestGetOneTask_NotFound(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks/9999")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var errorResponse map[string]string
	err = json.NewDecoder(resp.Body).Decode(&errorResponse)
	assert.NoError(t, err)
	assert.Contains(t, errorResponse["message"], "Failed to get task")
}

func TestEditTask_InvalidID(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	task := dto.Task{
		Text:      "Updated Task",
		Completed: true,
	}

	taskJSON, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, server.URL+"/tasks/edit/abc", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var errorResponse map[string]string
	err = json.NewDecoder(resp.Body).Decode(&errorResponse)
	assert.NoError(t, err)
	assert.Contains(t, errorResponse["message"], "Invalid request body")
}

func TestEditTask_InvalidJSON(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	invalidJSON := `{"text": 123, "completed": "not a boolean"}`

	req, err := http.NewRequest(http.MethodPut, server.URL+"/tasks/edit/1", strings.NewReader(invalidJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var errorResponse map[string]string
	err = json.NewDecoder(resp.Body).Decode(&errorResponse)
	assert.NoError(t, err)
	assert.Contains(t, errorResponse["message"], "Invalid request body")
}

func TestDeleteTask_InvalidID(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	req, err := http.NewRequest(http.MethodDelete, server.URL+"/tasks/delete/abc", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var errorResponse map[string]string
	err = json.NewDecoder(resp.Body).Decode(&errorResponse)
	assert.NoError(t, err)
	assert.Contains(t, errorResponse["message"], "Invalid request body")
}

func TestGetAllTasks_InvalidQueryParam(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks?completed=invalid")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var errorResponse map[string]string
	err = json.NewDecoder(resp.Body).Decode(&errorResponse)
	assert.NoError(t, err)
	assert.Contains(t, errorResponse["message"], "Failed to get tasks")
}
