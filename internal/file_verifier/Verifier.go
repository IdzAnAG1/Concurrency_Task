package file_verifier

import (
	"concurrency_task/internal/channels"
	"concurrency_task/internal/config"
	"concurrency_task/internal/file_verifier/change_detector"
	"concurrency_task/internal/file_verifier/errors_handler"
	"concurrency_task/internal/file_verifier/file_readiness_detector"
	"concurrency_task/internal/file_verifier/injection_of_function_init"
	"concurrency_task/internal/logger"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/file_handler"
	"time"
)

type Mechanisms interface {
	Launch(channels.Channel)
}
type Verifier struct {
	PathToMethodsDirectory string
	Interval               time.Duration
	QuanFilesInDirectory   int
	TCStorage              *task_code_storage.TCStorage
}

func NewVerifier(cfg config.Config) *Verifier {
	return &Verifier{
		PathToMethodsDirectory: cfg.PathToMethodsDirectory,
		Interval:               cfg.Interval,
		QuanFilesInDirectory:   0,
	}
}

func (v *Verifier) Run() error {
	logs, err := logger.NewLogger()
	if err != nil {
		return err
	}
	defer logs.Sync()

	channel := channels.NewChannel(*logs)
	Store := task_code_storage.NewTCStorage(*logs)
	err = Store.Initialize(v.PathToMethodsDirectory)
	if err != nil {
		return err
	}

	v.TCStorage = Store
	files, err := file_handler.GetFilesInDirectory(v.PathToMethodsDirectory)
	if err != nil {
		channel.SendErrorsToChannel(err)
	}
	v.QuanFilesInDirectory = len(files)

	Chad := change_detector.NewChad(
		*logs,
		v.PathToMethodsDirectory,
		v.Interval,
		v.QuanFilesInDirectory,
		v.TCStorage,
	)
	Fired := file_readiness_detector.NewFired(*logs, v.TCStorage)
	Infinit := injection_of_function_init.NewInfinit(*logs, v.PathToMethodsDirectory, v.TCStorage)
	errHan := errors_handler.NewErrorsHandler(*logs)

	mechanismsStart(channel, errHan, Chad, Fired, Infinit)
	for {
		select {
		case val := <-channel.ReadChangeFromChannel():
			{
				if val {

				}
			}
		}
	}
}

func mechanismsStart(channel *channels.Channel, mechanisms ...Mechanisms) {
	for _, el := range mechanisms {
		el.Launch(*channel)
	}
}
