package task_storage

import (
	"concurrency_task/internal/tasks"
	"concurrency_task/internal/utils/go_uuid"
	"fmt"
	"strings"
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
		_ = fmt.Errorf("The task storage has not been initialized ")
	}
	return StorageInstance
}
func (ts *TaskStorage) AddInStorage(name string, task tasks.ConcurrencyTask) {
	if strings.TrimSpace(name) == "" {
		name = go_uuid.Uid()
	}
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.Storage[name] = task
}

func (ts *TaskStorage) GetKeys() []string {
	arr := make([]string, 0)
	ts.mu.Lock()
	for key, _ := range ts.Storage {
		arr = append(arr, key)
	}
	ts.mu.Unlock()
	return arr
}

func (ts *TaskStorage) taskIsLocatedInTheRepository(taskName string) bool {
	ts.mu.Lock()
	_, ex := ts.Storage[taskName]
	ts.mu.Unlock()
	return ex
}
