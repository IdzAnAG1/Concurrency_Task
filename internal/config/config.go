package config

import (
	"concurrency_task/internal/tasks/task_code_storage"
	"time"
)

type Config struct {
	PathToMethodsDirectory string
	Interval               time.Duration
	QuanFilesInDirectory   int
	TCStorage              *task_code_storage.TCStorage
}
