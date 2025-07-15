package task_code_storage

import (
	"concurrency_task/internal/utils/file_handler"
	"errors"
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

func (ts *TCStorage) Put(filename string, code string) error {
	if filename == "" {
		return errors.New("filename cannot be empty")
	}
	ts.logger.Info(fmt.Sprintf("%s was updated in task_code_storage", filename))
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.Storage[filename] = code
	return nil
}
func (ts *TCStorage) Get(filename string) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	return ts.Storage[filename], nil
}
func (ts *TCStorage) Delete(filename string) {
	ts.logger.Info(fmt.Sprintf("Stopping tracking the contents of the %s file ", filename))
	ts.mu.Lock()
	defer ts.mu.Unlock()
	delete(ts.Storage, filename)
}

func (ts *TCStorage) Initialize(pathToDir string) error {
	if pathToDir == "" {
		return errors.New("path to Directory cannot be empty")
	}
	Files, err := file_handler.GetFilesInDirectory(pathToDir)
	if err != nil {
		return err
	}
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

func (ts *TCStorage) Len() int {
	ts.mu.Lock()
	retVal := len(ts.Storage)
	ts.mu.Unlock()
	return retVal
}

func (ts *TCStorage) GetKeys() []string {
	var keys = make([]string, ts.Len())
	ts.mu.Lock()
	for key, _ := range ts.Storage {
		keys = append(keys, key)
	}
	ts.mu.Unlock()
	return keys
}
