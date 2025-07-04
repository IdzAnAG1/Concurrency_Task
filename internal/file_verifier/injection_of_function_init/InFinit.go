package injection_of_function_init

import (
	"concurrency_task/internal/models"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/variables"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Infinit struct {
	PathToDir   string
	CodeStorage *task_code_storage.TCStorage
}

func NewInfinit(path string, storage *task_code_storage.TCStorage) *Infinit {
	return &Infinit{
		path, storage,
	}
}
func (i *Infinit) Launch(channels models.Channel) {
	go func() {
		for {
			select {
			case ind := <-channels.ContentIndicator:
				fmt.Printf("%+v", ind)
				i.userStructIsNotExist(ind)
			}
		}
	}()
}

func (i *Infinit) userStructIsNotExist(FiredMSG *models.InfinitData_v2) {
	str := i.CodeStorage.Get(FiredMSG.FileName)
	lines := strings.Split(str, "\n")
	index := FiredMSG.Indicator.FileFullness[variables.USER_STRUCT] - 1
	lines[index] = "//You should write user struct which is needed to implement the interface"
	str = strings.Join(lines, "\n")
	err := os.WriteFile(filepath.Join(i.PathToDir, FiredMSG.FileName), []byte(str), 0644)
	if err != nil {

	}
}
