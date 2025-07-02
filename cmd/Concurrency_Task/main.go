package main

import (
	"concurrency_task/internal/file_verifier/change_detector"
	"concurrency_task/internal/file_verifier/file_readiness_detector"
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

	Fired.Launch(channels.Content)
	Chad.Launch(channels.Boolean, channels.Content)

	for {
		select {
		case val := <-channels.Boolean:
			{
				if val {
					fmt.Println("channel is catch the signal")
				}
			}
		default:

		}
	}

}
