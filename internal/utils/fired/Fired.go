package fired

import (
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/variables"
	"regexp"
	"strings"
)

// Fired - Files Readiness Detector
type Fired struct {
	tcStorage *task_code_storage.TCStorage
}

func NewFired() *Fired {
	return &Fired{}
}

func (f *Fired) Launch(channel chan string) {
	go func() {
		for {
			select {
			case val := <-channel:
				{
					if fileIsReadyImp(val) {

					}
				}
			}
		}
	}()
}

func fileIsReadyImp(content string) bool {
	lines := strings.Split(content, "\n")
	if ok := fileHasContain(variables.REG_EXP_USERSTUCT, lines); !ok {
		return false
	}
	if ok := fileHasContain(variables.REG_EXP_FUNCINIT, lines); !ok {
		return false
	}
	return true
}

func fileHasContain(regularExpression string, content []string) bool {
	r := regexp.MustCompile(regularExpression)
	for _, el := range content {
		if r.MatchString(el) {
			return true
		}
	}
	return false
}
