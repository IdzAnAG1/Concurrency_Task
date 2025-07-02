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

func (f *Fired) Launch(channel chan string) {
	go func() {
		for {
			select {
			case val := <-channel:
				{
					indicator := f.fileIsReadyImp(val)
					fmt.Println(indicator)
				}
			}
		}
	}()
}

func (f *Fired) fileIsReadyImp(content string) (indicator models.ReadinessIndicator) {
	code := f.tcStorage.Get(content)
	lines := strings.Split(code, "\n")
	for key, value := range variables.RagExpressions {
		flag := regex.Contains(value, lines)
		switch key {
		case variables.USER_STRUCT:
			indicator.UserStructIsExist = flag

		case variables.IMPLEMENTED_FUNC:
			indicator.InterfaceImplementationIsExist = flag

		case variables.FUNC_INIT:
			indicator.InitFuncIsExist = flag
		}
	}
	return
}
