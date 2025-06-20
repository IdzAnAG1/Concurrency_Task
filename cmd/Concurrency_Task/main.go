package main

import (
	_ "concurrency_task/internal/tasks/task_impl"
	"concurrency_task/internal/tasks/task_storage"
	"concurrency_task/internal/utils/capabilityChecker"
	"fmt"
	"time"
)

func main() {
	var (
		BooleanChannel = make(chan bool)
	)

	Tasks := task_storage.GetStorageInstance()

	go func(store task_storage.TaskStorage, booleanChannel chan bool) {
		CapChecker := capabilityChecker.NewCapChecker("internal/tasks/task_impl", 500*time.Millisecond)
		fmt.Println("Cap Checker Started")
		CapChecker.LaunchChecker(booleanChannel)
	}(*Tasks, BooleanChannel)

	for {
		select {
		case val := <-BooleanChannel:
			{
				if val {
					fmt.Println("New task")
					time.Sleep(5 * time.Second)
				}
			}
		default:
			fmt.Println("Ничего нового")
		}
	}
}
