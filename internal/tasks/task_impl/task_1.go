package task_impl

import (
	"concurrency_task/internal/tasks/task_storage"
	"fmt"
)
//You should write user struct which is needed to implement the interface
type Task1 struct {
	Name string
}

func (t Task1) Launch() {
	fmt.Println("Launching test test_2 Test_3 -Test_4- test_6 test_7")
}

func init() {
	Storage := task_storage.GetStorageInstance()
	Storage.AddInStorage("task_1", &Task1{})
}
