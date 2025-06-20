package capabilityChecker

import (
	"fmt"
	"log"
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
		quanFilesInDirectory:   getQuanFilesInDirectory(pathToMethodsDirectory),
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
			if val, ok := cc.isFileWasUpdated(); ok == true {
				fmt.Println(fmt.Sprintf("%s was updated", val))
			}
		}
	}
}

func getQuanFilesInDirectory(path string) int {
	FilesIntoDirectory, err := os.ReadDir(path + "/")
	if err != nil {
		log.Fatal(err)
	}
	return len(FilesIntoDirectory)
}

func (cc *CapChecker) isDirectoryWasUpdated() bool {
	filesNow := getQuanFilesInDirectory(cc.PathToMethodsDirectory)
	if filesNow != cc.quanFilesInDirectory {
		cc.quanFilesInDirectory = filesNow
		return true
	}
	return false
}

func (cc *CapChecker) isFileWasUpdated() (string, bool) {
	return "", false
}
