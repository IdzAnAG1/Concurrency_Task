package file_readiness_detector

import (
	"concurrency_task/internal/channels"
	"concurrency_task/internal/models"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/regex"
	"concurrency_task/internal/variables"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

// Fired - Files Readiness Detector
type Fired struct {
	logger    zap.Logger
	tcStorage *task_code_storage.TCStorage
}

func NewFired(logger zap.Logger, store *task_code_storage.TCStorage) *Fired {
	return &Fired{
		logger,
		store,
	}
}

func (f *Fired) Launch(channels channels.Channel) {
	go func() {
		for {
			select {
			case val := <-channels.ReadContentFromChannel():
				{
					test := f.fileIsReadyImp(val)
					channels.SendToChannelContentIndicator(test)
				}
			}
		}
	}()
}

func (f *Fired) fileIsReadyImp(content string) *models.InfinitData {
	code := f.tcStorage.Get(content)
	lines := strings.Split(code, "\n")
	infData := &models.InfinitData{
		FileName:  content,
		Indicator: *models.NewReadinessIndicator(),
	}
	for key, val := range variables.RegExpressions {
		index, flag := regex.Contains(val, lines)
		f.logger.Info(fmt.Sprintf("The %s is present: %t", key, flag))
		infData.Indicator.Put(key, index)
	}

	return infData
}
