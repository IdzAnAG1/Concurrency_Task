package task_storage

import (
	"concurrency_task/internal/tasks"
	"sync"
)

var (
	StorageInstance *TaskStorage
)

type TaskStorage struct {
	mu      sync.Mutex
	Storage map[string]tasks.ConcurrencyTask
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		mu:      sync.Mutex{},
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
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.Storage[name] = task
}
