package task_impl

import (
	"concurrency_task/internal/tasks/task_storage"
	"fmt"
)

// Come up with your own name for the structure and
// replace the current UserStruct_c67eb694_6a0c_44e4_9e71_4b6db28718f4  structure name
type Test struct {
	name string
	age  int
}

// This is an interface implementation function.
// To run the program, it uses a receiver with a name created automatically.
// Change the name in the receiver to the one you created.
func (t *Test) Launch() {
	fmt.Println("Fix")
}

// This function allows you to immediately add your task to the execution stream.
// Change it only if you change the name of the structure created automatically.
func init() {
	Storage := task_storage.GetStorageInstance()
	Storage.AddInStorage("Task_Test",
		&Test{})
}
