package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rahulp18/todo/db"
	"github.com/rahulp18/todo/handler"
	"github.com/rahulp18/todo/service"
	"github.com/rahulp18/todo/store"
)

func main() {
	// Initial memory for storing
	// memoryStore := store.NewMemoryTaskStore()

	//   DB CONNECTION
	dbConn := db.NewPostgres()
	defer dbConn.Close()

	// fileStore := store.NewFileTaskStore("tasks.json")
	dbStore := store.NewPgTaskStore(dbConn)
	taskServices := service.NewTaskService(dbStore)

	handler.SetTaskService(taskServices)
	http.HandleFunc("/tasks", handler.Tasks)
	http.HandleFunc("/tasks/", handler.TaskById)

	fmt.Println("Server is listing ar port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}

}
