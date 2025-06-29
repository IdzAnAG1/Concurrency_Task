package fired

import (
	"concurrency_task/internal/variables"
	"fmt"
	"regexp"
	"strings"
)

// Fired - Files Readiness Detector
type Fired struct {
}

func NewFired() *Fired {
	return &Fired{}
}

func (f *Fired) Run() {
	go func() {

	}()
}

func (f *Fired) FileIsReadyImp(contentsChannel chan string) bool {
	str := strings.Split(<-contentsChannel, " ")
	fmt.Println(str)
	return true
}

func fileHasUserStruct(content []string) bool {
	for _, el := range content {
		if elementIsValid(variables.REG_EXP_USERSTUCT, el) {
			return true
		}
	}
	return false
}

func fileHasFunctionForImplement(content []string) bool {
	for _, el := range content {
		if elementIsValid(variables.REG_EXP_FUNCINIT, el) {
			return true
		}
	}
	return false
}
func elementIsValid(regularExpression, line string) bool {
	r := regexp.MustCompile(regularExpression)
	return r.MatchString(line)
}
