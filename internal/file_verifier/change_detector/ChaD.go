package change_detector

import (
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
	quanFilesInDirectory   int
	tcStorage              *task_code_storage.TCStorage
}

func NewChad(pathToMethodsDirectory string, interval time.Duration, filesNumber int, storage *task_code_storage.TCStorage) *ChaD {
	return &ChaD{
		PathToMethodsDirectory: pathToMethodsDirectory,
		Interval:               interval,
		quanFilesInDirectory:   filesNumber,
		tcStorage:              storage,
	}
}

func (ch *ChaD) Launch(channel chan bool, channelToFired chan string) {
	go func() {
		ch.tcStorage.Initialize(ch.PathToMethodsDirectory)
		ticker := time.NewTicker(ch.Interval)
		defer ticker.Stop()

		for {
			files := file_handler.GetFilesInDirectory(ch.PathToMethodsDirectory)
			select {
			case <-ticker.C:
				if ch.isDirectoryWasUpdated(len(files)) {
					channel <- true
				}
				if val, ok := ch.isFileWasUpdated(files); ok {
					channelToFired <- val
					channel <- true
				}
			}
		}
	}()
}

func (ch *ChaD) isDirectoryWasUpdated(filesNow int) bool {
	if filesNow != ch.quanFilesInDirectory {
		ch.quanFilesInDirectory = filesNow
		return true
	}
	return false
}

func (ch *ChaD) isFileWasUpdated(filesInDir []os.DirEntry) (string, bool) {
	for _, file := range filesInDir {
		currentCode := file_handler.ReadFromFile(ch.PathToMethodsDirectory, file)
		if ch.isCurrentContentNotActual(currentCode, file.Name()) {
			return file.Name(), true
		}
	}
	return "", false
}

func (ch *ChaD) isCurrentContentNotActual(currentContent, filename string) bool {
	savedEntry := hash.ConvertToHash(ch.tcStorage.Get(filename))
	if hash.ConvertToHash(currentContent) != savedEntry {
		ch.tcStorage.Put(filename, currentContent)
		return true
	}
	return false
}
