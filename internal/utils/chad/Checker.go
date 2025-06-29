package chad

import (
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/general"

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
		quanFilesInDirectory:   filesNumber,
		tcStorage:              task_code_storage.NewTCStorage(),
	}
}

// _________________________________________________________________________

func (ch *ChaD) Launch(channel chan bool) {
	ch.tcStorage.Initialize(ch.PathToMethodsDirectory)
	ticker := time.NewTicker(ch.Interval)
	defer ticker.Stop()

	for {
		files := general.GetFilesInDirectory(ch.PathToMethodsDirectory)
		select {
		case <-ticker.C:
			if ch.isDirectoryWasUpdated(len(files)) {
				channel <- true
			}
			if _, ok := ch.isFileWasUpdated(files); ok {
				channel <- true
			}
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
		currentCode := general.ReadFromFile(ch.PathToMethodsDirectory, file)
		hashCode := general.ConvertToHash(currentCode)
		if currentCode != ch.tcStorage.Storage[file.Name()] {
			ch.tcStorage.Storage[file.Name()] = hashCode
			return file.Name(), true
		}
	}
	return "", false
}

/*
func (ch *ChaD) isFileWasUpdated_v2_MayBe(filesInDir []os.DirEntry) string {
	updatedFiles := make(chan string)
	wg := sync.WaitGroup{}
	for _, file := range filesInDir {
		wg.Add(1)
		go func(f os.DirEntry, chann chan string) {
			defer wg.Done()
			ch.mu.Lock()
			currentCode := general.ReadFromFile(ch.PathToMethodsDirectory, f)
			hashCode := general.ConvertToHash(currentCode)
			if hashCode != ch.tcStorage.Storage[f.Name()] {
				ch.tcStorage.Storage[f.Name()] = hashCode
				updatedFiles <- f.Name()
			}
			ch.mu.Unlock()
		}(file, updatedFiles)
	}
	wg.Wait()
	return <-updatedFiles
}
*/
