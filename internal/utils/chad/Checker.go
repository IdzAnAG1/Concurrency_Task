package chad

import (
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/general"
	"fmt"
	"os"
	"sync"
	"time"
)

// Переделать в DRY

// ChaD - CHAnge Detector
type ChaD struct {
	mu                     sync.Mutex
	PathToMethodsDirectory string
	Interval               time.Duration
	quanFilesInDirectory   int
	tcStorage              *task_code_storage.TCStorage
}

func NewChad(pathToMethodsDirectory string, interval time.Duration, filesNumber int) *ChaD {
	return &ChaD{
		mu:                     sync.Mutex{},
		PathToMethodsDirectory: pathToMethodsDirectory,
		Interval:               interval,
		quanFilesInDirectory:   filesNumber, // --
		tcStorage:              task_code_storage.NewTCStorage(),
	}
}

// _________________________________________________________________________

func (ch *ChaD) LaunchChad(channel chan bool) {
	ticker := time.NewTicker(ch.Interval)
	defer ticker.Stop()

	for {
		files := general.GetFilesInDirectory(ch.PathToMethodsDirectory)
		ch.tcStorage.Update(files, ch.PathToMethodsDirectory)
		select {
		case <-ticker.C:

		}
	}
}

// _________________________________________________________________________
func (ch *ChaD) isDirectoryWasUpdated(filesNow int) bool {
	if filesNow != ch.quanFilesInDirectory {
		ch.quanFilesInDirectory = filesNow
		return true
	}
	return false
}

func (ch *ChaD) isFileWasUpdated(filesInDir []os.DirEntry) (string, bool) {
	for _, file := range filesInDir {
		currentCode := general.ConvertToHash(general.ReadFromFile(ch.PathToMethodsDirectory, file))
		if currentCode != ch.tcStorage.Storage[file.Name()] {
			ch.tcStorage.Storage[file.Name()] = currentCode
			return fmt.Sprintf("%s - %s", time.Now().Format("02-01-2006 15:04:05"), file.Name()), true
		}
	}
	return "", false
}

func isFileWasUpdated_v2() {
	// конкурентно решить проблему

}
