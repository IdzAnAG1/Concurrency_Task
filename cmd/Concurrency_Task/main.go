package main

import (
	"concurrency_task/internal/config"
	"concurrency_task/internal/file_verifier"
	_ "concurrency_task/internal/tasks/task_impl"
	"time"
)

func main() {
	Cfg := config.Config{
		PathToMethodsDirectory: "internal/tasks/task_impl",
		Interval:               250 * time.Millisecond,
	}
	v := file_verifier.NewVerifier(Cfg)
	v.Run()
}
