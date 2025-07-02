package task_code_storage

import (
	"concurrency_task/internal/utils/file_handler"
	"sync"
)

type TCStorage struct {
	mu      sync.Mutex
	Storage map[string]string
}

func NewTCStorage() *TCStorage {
	return &TCStorage{
		mu:      sync.Mutex{},
		Storage: make(map[string]string),
	}
}

func (ts *TCStorage) Put(filename string, code string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.Storage[filename] = code
}
func (ts *TCStorage) Get(filename string) string {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	return ts.Storage[filename]
}
func (ts *TCStorage) Delete(filename string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	delete(ts.Storage, filename)
}

func (ts *TCStorage) Initialize(pathToDir string) {
	Files := file_handler.GetFilesInDirectory(pathToDir)
	for _, file := range Files {
		fileContent := file_handler.ReadFromFile(pathToDir, file)
		ts.mu.Lock()
		if _, exists := ts.Storage[file.Name()]; !exists {
			ts.Storage[file.Name()] = fileContent
		}
		ts.mu.Unlock()
	}
}
