package tasks

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

var (
	tasks    []Task
	fileName = "tasks.json"
)

func LoadTask() error {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil
	}

	return json.Unmarshal(file, &tasks)
}

func SaveTasks() error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return nil
	}

	return os.WriteFile(fileName, data, 0644)
}

func AddTask(title string) error {
	task := Task{
		ID:       len(tasks) + 1,
		Title:    title,
		Complete: false,
	}

	tasks = append(tasks, task)
	if err := SaveTasks(); err != nil {
		return err
	}
	fmt.Println("Task added:", title)
	return nil
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No Task Found")
		return
	}

	fmt.Println("To-Do List:")

	for _, task := range tasks {
		status := "X"
		if task.Complete {
			status = "V"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
	}
}
func CompleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			if tasks[i].Complete {
				return fmt.Errorf("task %d is already completed", id)
			}
			tasks[i].Complete = true
			if err := SaveTasks(); err != nil {
				return err
			}
			fmt.Println("Task completed:", task.Title)
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}
