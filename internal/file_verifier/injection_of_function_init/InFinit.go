package injection_of_function_init

import (
	"concurrency_task/internal/channels"
	"concurrency_task/internal/models"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/go_uuid"
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

func NewInfinit(pathToDir string, storage *task_code_storage.TCStorage) *Infinit {
	return &Infinit{
		PathToDir:   pathToDir,
		CodeStorage: storage,
	}
}
func (i *Infinit) Launch(channels channels.Channel) {
	go func() {
		for {
			select {
			case ind := <-channels.ReadInfDataFromChannel():
				i.userStructIsNotExist(ind)
			}
		}
	}()
}

func (i *Infinit) userStructIsNotExist(FiredMSG *models.InfinitData) {
	str := i.CodeStorage.Get(FiredMSG.FileName)
	temp := ""
	linesArray := strings.Split(str, "\n")
	if FiredMSG.Indicator.FileFullness[variables.USER_STRUCT] == -1 {
		linesArray, temp = writeUserStructPattern(linesArray)
	}
	if FiredMSG.Indicator.FileFullness[variables.IMPLEMENTED_FUNC] == -1 {
		linesArray = writeFunctionForImplementation(temp, linesArray)
	}
	if FiredMSG.Indicator.FileFullness[variables.FUNC_INIT] == -1 {
		linesArray = writeFuncInit(temp, linesArray)
	}
	str = strings.Join(linesArray, "\n")
	err := os.WriteFile(filepath.Join(i.PathToDir, FiredMSG.FileName), []byte(str), 0644)
	if err != nil {

	}
}

func writeUserStructPattern(StrArr []string) ([]string, string) {
	temp := fmt.Sprintf(variables.UserStructStartName, strings.ReplaceAll(go_uuid.Uid(), "-", "_"))
	line := fmt.Sprintf(variables.UserStructText, temp)
	comment := fmt.Sprintf(variables.CommentForUserStruct, temp)
	StrArr = append(StrArr, comment, line)
	return StrArr, temp
}

func writeFunctionForImplementation(structureName string, strArr []string) []string {
	text := fmt.Sprintf(variables.FunctionForImplementation, structureName)
	strArr = append(strArr, variables.CommentForFunctionForImplementation, text)
	return strArr
}

func writeFuncInit(structName string, strArr []string) []string {
	text := fmt.Sprintf(variables.FuncInitText, structName, structName)
	strArr = append(strArr, variables.CommentForFunctionInit, text)
	return strArr
}
