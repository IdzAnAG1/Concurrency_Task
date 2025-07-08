package file_readiness_detector

import (
	"concurrency_task/internal/config"
	"concurrency_task/internal/models"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/regex"
	"concurrency_task/internal/variables"
	"fmt"
	"strings"
)

// Fired - Files Readiness Detector
type Fired struct {
	tcStorage *task_code_storage.TCStorage
}

func NewFired(cfg *config.Config) *Fired {
	return &Fired{cfg.TCStorage}
}

func (f *Fired) Launch(channels models.Channel) {
	go func() {
		for {
			select {
			case val := <-channels.Content:
				{
					fmt.Println("Cathch in fired")
					test := f.fileIsReadyImp(val)
					channels.ContentIndicator <- test
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
		fmt.Println("!~|~!", flag, index, "!~|~!")
		infData.Indicator.Put(key, index)
	}

	return infData
}
