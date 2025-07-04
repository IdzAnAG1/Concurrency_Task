package main

import (
	"concurrency_task/internal/file_verifier/change_detector"
	"concurrency_task/internal/file_verifier/file_readiness_detector"
	"concurrency_task/internal/file_verifier/injection_of_function_init"
	"concurrency_task/internal/models"
	"concurrency_task/internal/tasks/task_code_storage"
	_ "concurrency_task/internal/tasks/task_impl"
	"concurrency_task/internal/utils/file_handler"
	"fmt"
	"time"
)

func main() {
	channels := models.NewChannel()
	files := file_handler.GetFilesInDirectory("internal/tasks/task_impl")
	store := task_code_storage.NewTCStorage()
	Fired := file_readiness_detector.NewFired(store)
	Chad := change_detector.NewChad(
		"internal/tasks/task_impl",
		time.Millisecond*250,
		len(files), store)
	Inf := injection_of_function_init.NewInfinit(Chad.PathToMethodsDirectory, store)
	Inf.Launch(*channels)
	Fired.Launch(*channels)
	Chad.Launch(channels.ContentChange, channels.Content)

	for {
		select {
		case val := <-channels.ContentChange:
			{
				if val {
					fmt.Println("channel is catch the signal")
				}
			}
		default:

		}
	}
}
