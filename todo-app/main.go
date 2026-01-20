package main

import (
	"fmt"
	"net/http"

	"github.com/rahulp18/todo/handler"
	"github.com/rahulp18/todo/service"
	"github.com/rahulp18/todo/store"
)

func main() {
	// Initial memory for storing
	// memoryStore := store.NewMemoryTaskStore()
	fileStore := store.NewFileTaskStore("tasks.json")
	taskServices := service.NewTaskService(fileStore)

	handler.SetTaskService(taskServices)
	http.HandleFunc("/tasks", handler.Tasks)
	http.HandleFunc("/tasks/", handler.TaskById)

	fmt.Println("Server is listing ar port 8080")
	http.ListenAndServe(":8080", nil)

}
