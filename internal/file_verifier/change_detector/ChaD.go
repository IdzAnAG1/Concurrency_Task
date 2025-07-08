package change_detector

import (
	"concurrency_task/internal/config"
	"concurrency_task/internal/models"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/file_handler"
	"concurrency_task/internal/utils/hash"

	"os"
	"time"
)

// ChaD - Change Detector
type ChaD struct {
	PathToMethodsDirectory string
	Interval               time.Duration
	QuanFilesInDirectory   int
	TCStorage              *task_code_storage.TCStorage
}

func NewChad(cfg *config.Config) *ChaD {
	return &ChaD{
		PathToMethodsDirectory: cfg.PathToMethodsDirectory,
		Interval:               cfg.Interval,
		QuanFilesInDirectory:   cfg.QuanFilesInDirectory,
		TCStorage:              cfg.TCStorage,
	}
}

func (ch *ChaD) Launch(channel models.Channel) {
	go func() {
		ticker := time.NewTicker(ch.Interval)
		defer ticker.Stop()

		for {
			files := file_handler.GetFilesInDirectory(ch.PathToMethodsDirectory)
			select {
			case <-ticker.C:
				if ch.isDirectoryWasUpdated(len(files)) {
					channel.ContentChange <- true
				}
				nameFile, err := ch.isFileWasUpdated(files)
				if err != nil {
					channel.Errors <- err
				}
				if nameFile != "" {
					channel.Content <- nameFile
				}
			}
		}
	}()
}

func (ch *ChaD) isDirectoryWasUpdated(filesNow int) bool {
	if filesNow != ch.QuanFilesInDirectory {
		ch.QuanFilesInDirectory = filesNow
		return true
	}
	return false
}

func (ch *ChaD) isFileWasUpdated(filesInDir []os.DirEntry) (string, error) {
	for _, file := range filesInDir {
		currentCode, err := file_handler.ReadFromFile(ch.PathToMethodsDirectory, file)
		if err != nil {
			return "", err
		}
		if ch.isCurrentContentNotActual(currentCode, file.Name()) {
			return file.Name(), nil
		}
	}
	return "", nil
}

func (ch *ChaD) isCurrentContentNotActual(currentContent, filename string) bool {
	savedEntry := hash.ConvertToHash(ch.TCStorage.Get(filename))
	if hash.ConvertToHash(currentContent) != savedEntry {
		ch.TCStorage.Put(filename, currentContent)
		return true
	}
	return false
}
