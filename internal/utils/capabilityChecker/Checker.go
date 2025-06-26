package capabilityChecker

import (
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/general"
	"fmt"
	"os"
	"sync"
	"time"
)

type CapChecker struct {
	mu                     sync.Mutex
	PathToMethodsDirectory string
	Interval               time.Duration
	quanFilesInDirectory   int
	tcStorage              *task_code_storage.TCStorage
}

func NewCapChecker(pathToMethodsDirectory string, interval time.Duration) *CapChecker {
	return &CapChecker{
		mu:                     sync.Mutex{},
		PathToMethodsDirectory: pathToMethodsDirectory,
		Interval:               interval,
		quanFilesInDirectory:   len(general.GetFilesInDirectory(pathToMethodsDirectory)),
		tcStorage:              task_code_storage.NewTCStorage(),
	}
}

func (cc *CapChecker) LaunchChecker(channel chan bool) {
	cc.tcStorage.Initialization(general.GetFilesInDirectory(cc.PathToMethodsDirectory), cc.PathToMethodsDirectory)
	fmt.Println(cc.tcStorage)
	ticker := time.NewTicker(cc.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if cc.isDirectoryWasUpdated() {
				fmt.Println("Add new file in the directory")
				channel <- true
			}
			if val, ok := cc.isFileWasUpdated(general.GetFilesInDirectory(cc.PathToMethodsDirectory),
				cc.tcStorage); ok == true {
				fmt.Println(fmt.Sprintf("%s was updated", val))
				channel <- true
			}
		}
	}
}

func (cc *CapChecker) isDirectoryWasUpdated() bool {
	filesNow := len(general.GetFilesInDirectory(cc.PathToMethodsDirectory))
	if filesNow != cc.quanFilesInDirectory {
		cc.quanFilesInDirectory = filesNow
		return true
	}
	return false
}

func (cc *CapChecker) isFileWasUpdated(filesInDir []os.DirEntry, tcs *task_code_storage.TCStorage) (string, bool) {
	for _, file := range filesInDir {
		currentCode := general.ConvertToHash(general.ReadFromFile(cc.PathToMethodsDirectory, file))
		if currentCode != tcs.Storage[file.Name()] {
			tcs.Storage[file.Name()] = currentCode
			return fmt.Sprintf("%s - %s", time.Now().Format("02-01-2006 15:04:05"), file.Name()), true
		}
	}
	return "", false
}
