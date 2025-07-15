package file_readiness_detector

import (
	"concurrency_task/internal/channels"
	"concurrency_task/internal/models"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/regex"
	"concurrency_task/internal/variables"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"strings"
	"sync"
)

// Fired - Files Readiness Detector
type Fired struct {
	logger    zap.Logger
	channels  channels.Channel
	tcStorage *task_code_storage.TCStorage
}

func NewFired(logger zap.Logger, channels channels.Channel, store *task_code_storage.TCStorage) *Fired {
	return &Fired{
		logger,
		channels,
		store,
	}
}

func (f *Fired) Launch(ctx context.Context, group *sync.WaitGroup) {
	group.Add(1)
	go func() {
		defer group.Done()
		f.logger.Info("Files Readiness Detector was launched")
		for {
			select {
			case val := <-f.channels.ReadContentFromChannel():
				{
					test, err := f.fileIsReadyImp(val)
					if err != nil {
						f.channels.SendErrorsToChannel(err)
					}
					f.channels.SendToChannelContentIndicator(test)
				}
			case <-ctx.Done():
				f.logger.Info("The completion signal is received in Files Readiness Detector")
				return
			}
		}
	}()
}

func (f *Fired) fileIsReadyImp(content string) (*models.InfinitData, error) {
	code, err := f.tcStorage.Get(content)
	if err != nil {
		return nil, err
	}
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
	return infData, nil
}
