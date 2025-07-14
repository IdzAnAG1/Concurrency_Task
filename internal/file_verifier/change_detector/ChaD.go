package change_detector

import (
	"concurrency_task/internal/channels"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/file_handler"
	"concurrency_task/internal/utils/hash"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"os"
	"time"
)

// ChaD - Change Detector
type ChaD struct {
	logger                 zap.Logger
	PathToMethodsDirectory string
	Interval               time.Duration
	QuanFilesInDirectory   int
	TCStorage              *task_code_storage.TCStorage
}

func NewChad(logger zap.Logger, pathToDirectory string, interval time.Duration, quanFiles int, store *task_code_storage.TCStorage) *ChaD {
	return &ChaD{
		logger:                 logger,
		PathToMethodsDirectory: pathToDirectory,
		Interval:               interval,
		QuanFilesInDirectory:   quanFiles,
		TCStorage:              store,
	}
}

func (ch *ChaD) Launch(ctx context.Context, Channels channels.Channel) {
	go func() {
		ch.logger.Info("Change Detector was launched")
		ticker := time.NewTicker(ch.Interval)
		defer ticker.Stop()

		for {
			files, err := file_handler.GetFilesInDirectory(ch.PathToMethodsDirectory)
			if err != nil {
				Channels.SendErrorsToChannel(err)
			}
			select {
			case <-ticker.C:
				if ch.isDirectoryWasUpdated(len(files)) {
					Channels.SendToChangeChannel(true)
				}
				nameFile, err1 := ch.isFileWasUpdated(files)
				if err1 != nil {
					Channels.SendErrorsToChannel(err1)
				}
				if nameFile != "" {
					Channels.SendToContentChannel(nameFile)
				}
			case <-ctx.Done():
				ch.logger.Info("The completion signal is received in Change Detector")
				return
			}
		}
	}()
}

func (ch *ChaD) isDirectoryWasUpdated(filesNow int) bool {
	if filesNow != ch.QuanFilesInDirectory {
		ch.logger.Info("There have been changes in the directory")
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
		if isActual, err1 := ch.isCurrentContentNotActual(currentCode, file.Name()); err1 != nil {
			return "", err1
		} else if isActual == true {
			return file.Name(), nil
		}
	}
	return "", nil
}

func (ch *ChaD) isCurrentContentNotActual(currentContent, filename string) (bool, error) {
	savedContent, err1 := ch.TCStorage.Get(filename)
	if err1 != nil {
		return false, err1
	}
	savedEntry := hash.ConvertToHash(savedContent)
	if hash.ConvertToHash(currentContent) != savedEntry {
		ch.logger.Info(fmt.Sprintf("In the file %s updates have occurred", filename))
		ch.logger.Info(
			fmt.Sprintf(
				"The data in the repository does not match the contents of the %s, updating the contents of the repository",
				filename),
		)
		if err := ch.TCStorage.Put(filename, currentContent); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
