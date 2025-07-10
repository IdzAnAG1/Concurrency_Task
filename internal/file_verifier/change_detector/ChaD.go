package change_detector

import (
	"concurrency_task/internal/channels"
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

func NewChad(pathToDirectory string, interval time.Duration, quanFiles int, store *task_code_storage.TCStorage) *ChaD {
	return &ChaD{
		PathToMethodsDirectory: pathToDirectory,
		Interval:               interval,
		QuanFilesInDirectory:   quanFiles,
		TCStorage:              store,
	}
}

func (ch *ChaD) Launch(Channels channels.Channel) {
	go func() {
		ticker := time.NewTicker(ch.Interval)
		defer ticker.Stop()

		for {
			files := file_handler.GetFilesInDirectory(ch.PathToMethodsDirectory)
			select {
			case <-ticker.C:
				if ch.isDirectoryWasUpdated(len(files)) {
					Channels.SendToChangeChannel(true)
				}
				nameFile, err := ch.isFileWasUpdated(files)
				if err != nil {
					Channels.SendErrorsToChannel(err)
				}
				if nameFile != "" {
					Channels.SendToContentChannel(nameFile)
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
