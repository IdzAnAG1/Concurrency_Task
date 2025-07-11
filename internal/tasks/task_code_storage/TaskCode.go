package task_code_storage

import (
	"concurrency_task/internal/utils/file_handler"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

type TCStorage struct {
	logger  zap.Logger
	mu      sync.Mutex
	Storage map[string]string
}

func NewTCStorage(logger zap.Logger) *TCStorage {
	return &TCStorage{
		logger:  logger,
		mu:      sync.Mutex{},
		Storage: make(map[string]string),
	}
}

func (ts *TCStorage) Put(filename string, code string) {
	ts.logger.Info(fmt.Sprintf("%s was updated in task_code_storage", filename))
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
	ts.logger.Info(fmt.Sprintf("Stopping tracking the contents of the %s file ", filename))
	ts.mu.Lock()
	defer ts.mu.Unlock()
	delete(ts.Storage, filename)
}

func (ts *TCStorage) Initialize(pathToDir string) error {
	Files := file_handler.GetFilesInDirectory(pathToDir)
	for _, file := range Files {
		fileContent, err := file_handler.ReadFromFile(pathToDir, file)
		if err != nil {
			return err
		}
		ts.mu.Lock()
		if _, exists := ts.Storage[file.Name()]; !exists {
			ts.Storage[file.Name()] = fileContent
		}
		ts.mu.Unlock()
	}
	ts.logger.Info("The storage of the contents of the implementation files is initialized")
	return nil
}
