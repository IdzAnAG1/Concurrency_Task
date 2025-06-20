package task_storage

import (
	"concurrency_task/internal/tasks"
)

var (
	StorageInstance *TaskStorage
)

type TaskStorage struct {
	Storage map[string]tasks.ConcurrencyTask
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		Storage: make(map[string]tasks.ConcurrencyTask),
	}
}

func GetStorageInstance() *TaskStorage {
	if StorageInstance == nil {
		StorageInstance = NewTaskStorage()
	}
	return StorageInstance
}
func (ts *TaskStorage) AddInStorage(name string, task tasks.ConcurrencyTask) {
	ts.Storage[name] = task
}
