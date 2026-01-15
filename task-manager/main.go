package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rahulp18/task-manager/storage"
	"github.com/rahulp18/task-manager/task"
)

var (
	tasks []task.Task
	store storage.Storage
)

func main() {
	store = storage.FileStorage{Filename: "tasks.json"}
	loadedTasks, err := store.Load()
	if err != nil {
		fmt.Println("Error loading tasks", err)
		return
	}
	tasks = loadedTasks
	if len(os.Args) < 2 {
		fmt.Println("Please provide valid arguments")
		return
	}
	command := os.Args[1]

	switch command {
	case "add":
		addTask()
		store.Save(tasks)
	case "list":
		data, _ := store.Load()
		fmt.Println(data)
	case "drop":
		removeTask()
		store.Save(tasks)
	case "status":
		changeStatus()
		store.Save(tasks)
	default:
		fmt.Println("unknown command")
	}
}

func addTask() {
	if len(os.Args) < 3 {
		fmt.Println("Provide title")
		return
	}
	newTask := task.Task{
		ID:        len(tasks) + 1,
		Title:     os.Args[2],
		Status:    "pending",
		CreatedAt: time.Now().Local().String(),
	}
	tasks = append(tasks, newTask)
}
func removeTask() {
	if len(os.Args) < 3 {
		fmt.Println("Ops! we need task number to Drop it")
		return
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid task number")
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task removed", task.Title)
			return
		}
	}
	fmt.Println("Task not found")
}
func changeStatus() {
	if len(os.Args) < 4 {
		fmt.Println("Ufs we need 4 arguments")
		return
	}
	status := os.Args[3]
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid index found")
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			fmt.Println("Task status changed")
			return
		}
	}
	fmt.Println("Task not found")
}
