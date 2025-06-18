package task_storage

import "concurrency_task/internal/tasks"

type TaskStorage struct {
	Storage map[int]tasks.ConcurrencyTask
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		make(map[int]tasks.ConcurrencyTask),
	}
}

func (ts *TaskStorage) AddInStorage(number int, task tasks.ConcurrencyTask) {
	ts.Storage[number] = task
}
