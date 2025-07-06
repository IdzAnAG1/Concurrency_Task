package main

import (
	"concurrency_task/internal/file_verifier"
	_ "concurrency_task/internal/tasks/task_impl"
	"time"
)

func main() {
	v := file_verifier.NewVerifier(
		"internal/tasks/task_impl",
		250*time.Millisecond,
	)

	v.Run()
}
