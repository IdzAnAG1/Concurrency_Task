package capabilityChecker

import (
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
}

func NewCapChecker(pathToMethodsDirectory string, interval time.Duration) *CapChecker {
	return &CapChecker{
		mu:                     sync.Mutex{},
		PathToMethodsDirectory: pathToMethodsDirectory,
		Interval:               interval,
		quanFilesInDirectory:   len(general.GetFilesInDirectory(pathToMethodsDirectory)),
	}
}

func (cc *CapChecker) LaunchChecker(channel chan bool) {

	ticker := time.NewTicker(cc.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if cc.isDirectoryWasUpdated() {
				fmt.Println("Add new file in the directory")
				channel <- true
			}
			if val, ok := cc.isFileWasUpdated(general.GetFilesInDirectory(cc.PathToMethodsDirectory)); ok == true {
				fmt.Println(fmt.Sprintf("%s was updated", val))
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

func (cc *CapChecker) isFileWasUpdated(files []os.DirEntry) (string, bool) {
	wg := sync.WaitGroup{}
	for _, file := range files {
		wg.Add(1)
		go func(f os.DirEntry) {

		}(file)
	}
	return "", false
}
