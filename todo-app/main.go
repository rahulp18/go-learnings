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

	userStore := store.NewPostgresStore(dbConn)
	authService := service.NewAuthService(userStore)
	handler.SetAuthService(authService)

	http.HandleFunc("/tasks", handler.Tasks)
	http.HandleFunc("/tasks/", handler.TaskById)

	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)

	fmt.Println("Server is listing ar port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}

}
