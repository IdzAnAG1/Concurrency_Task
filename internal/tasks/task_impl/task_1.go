package task_impl

import (
	"concurrency_task/internal/tasks/task_storage"
	"fmt"
)

type Task1 struct {
	Name string
}

func (t Task1) Launch() {
	fmt.Println("Launching task")
}

func init() {
	Storage := task_storage.GetStorageInstance()
	Storage.AddInStorage("task_1", &Task1{})
}
