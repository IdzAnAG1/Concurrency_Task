package file_verifier

import (
	"concurrency_task/internal/channels"
	"concurrency_task/internal/config"
	"concurrency_task/internal/file_verifier/change_detector"
	"concurrency_task/internal/file_verifier/file_readiness_detector"
	"concurrency_task/internal/file_verifier/injection_of_function_init"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/file_handler"
	"fmt"
	"log"
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
		PathToMethodsDirectory: cfg.FileVerifier.PathToMethodsDirectory,
		Interval:               cfg.FileVerifier.Interval,
		QuanFilesInDirectory:   0,
	}
}

func (v *Verifier) Run() {
	channel := channels.NewChannel()
	Store := task_code_storage.NewTCStorage()
	err := Store.Initialize(v.PathToMethodsDirectory)
	if err != nil {

	}

	v.TCStorage = Store
	v.QuanFilesInDirectory = len(file_handler.GetFilesInDirectory(v.PathToMethodsDirectory))

	Chad := change_detector.NewChad(v.PathToMethodsDirectory, v.Interval, v.QuanFilesInDirectory, v.TCStorage)
	Fired := file_readiness_detector.NewFired(v.TCStorage)
	Infinit := injection_of_function_init.NewInfinit(v.PathToMethodsDirectory, v.TCStorage)

	mechanismsStart(channel, Chad, Fired, Infinit)
	for {
		select {
		case val := <-channel.ReadChangeFromChannel():
			{
				if val {
					fmt.Println("channel is catch the signal")
				}
			}
		case err := <-channel.ReadErrorsFromChannel():
			{
				if err != nil {
					log.Println(err)
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
