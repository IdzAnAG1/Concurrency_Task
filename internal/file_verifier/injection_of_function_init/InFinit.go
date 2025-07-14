package injection_of_function_init

import (
	"concurrency_task/internal/channels"
	"concurrency_task/internal/models"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/go_uuid"
	"concurrency_task/internal/variables"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"os"
	"path/filepath"
	"strings"
)

type Infinit struct {
	logger      zap.Logger
	PathToDir   string
	CodeStorage *task_code_storage.TCStorage
}

func NewInfinit(logger zap.Logger, pathToDir string, storage *task_code_storage.TCStorage) *Infinit {
	return &Infinit{
		logger:      logger,
		PathToDir:   pathToDir,
		CodeStorage: storage,
	}
}
func (i *Infinit) Launch(ctx context.Context, channels channels.Channel) {
	go func() {
		i.logger.Info("Injection function init was launched")

		for {
			select {
			case ind := <-channels.ReadInfDataFromChannel():
				if err := i.userStructIsNotExist(ind); err != nil {
					channels.SendErrorsToChannel(err)
				}
			case <-ctx.Done():
				i.logger.Info("The completion signal is received in Injection function init")
				return
			}
		}
	}()
}

func (i *Infinit) userStructIsNotExist(FiredMSG *models.InfinitData) error {
	str, err := i.CodeStorage.Get(FiredMSG.FileName)
	if err != nil {
		return err
	}
	temp := ""
	linesArray := strings.Split(str, "\n")
	if FiredMSG.Indicator.FileFullness[variables.USER_STRUCT] == -1 {
		i.logger.Info("Custom structure added")
		linesArray, temp = writeUserStructPattern(linesArray)
	}
	if FiredMSG.Indicator.FileFullness[variables.IMPLEMENTED_FUNC] == -1 {
		i.logger.Info("The interface implementation feature has been added")
		linesArray = writeFunctionForImplementation(temp, linesArray)
	}
	if FiredMSG.Indicator.FileFullness[variables.FUNC_INIT] == -1 {
		i.logger.Info("The function of automatically adding to ready-to-run functions has been added")
		linesArray = writeFuncInit(temp, linesArray)
	}
	str = strings.Join(linesArray, "\n")
	err = os.WriteFile(filepath.Join(i.PathToDir, FiredMSG.FileName), []byte(str), 0644)
	if err != nil {
		return err
	}
	return nil
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
