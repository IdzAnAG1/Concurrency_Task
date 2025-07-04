package file_readiness_detector

import (
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

func NewFired(store *task_code_storage.TCStorage) *Fired {
	return &Fired{store}
}

func (f *Fired) Launch(channels models.Channel) {
	go func() {
		for {
			select {
			case val := <-channels.Content:
				{
					fmt.Println("Cathch in fired")
					test := f.fileIsReadyImp_v2(val)
					channels.ContentIndicator <- test
				}
			}
		}
	}()
}

func (f *Fired) fileIsReadyImp_v2(content string) *models.InfinitData_v2 {
	code := f.tcStorage.Get(content)
	lines := strings.Split(code, "\n")

	infData := &models.InfinitData_v2{
		FileName:  content,
		Indicator: *models.NewReadinessIndicator(),
	}

	for key, val := range variables.RegExpressions {
		index, flag := regex.Contains(val, lines)
		if flag {
			infData.Indicator.Put(key, index)
		}
	}

	return infData
}
