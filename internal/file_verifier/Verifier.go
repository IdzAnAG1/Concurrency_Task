package file_verifier

import (
	"concurrency_task/internal/config"
	"concurrency_task/internal/file_verifier/change_detector"
	"concurrency_task/internal/file_verifier/file_readiness_detector"
	"concurrency_task/internal/file_verifier/injection_of_function_init"
	"concurrency_task/internal/models"
	"concurrency_task/internal/tasks/task_code_storage"
	"concurrency_task/internal/utils/file_handler"
	"fmt"
	"log"
	"time"
)

type Mechanisms interface {
	Launch(models.Channel)
}
type Verifier struct {
	Cfg config.Config
}

func NewVerifier(PathToMethodsDirectory string, Interval time.Duration) *Verifier {
	return &Verifier{
		config.Config{
			PathToMethodsDirectory: PathToMethodsDirectory,
			Interval:               Interval,
			QuanFilesInDirectory:   0,
			TCStorage:              nil,
		},
	}
}

func (v *Verifier) Run() {
	channel := models.NewChannel()
	Store := task_code_storage.NewTCStorage()
	err := Store.Initialize(v.Cfg.PathToMethodsDirectory)
	if err != nil {

	}

	v.Cfg.TCStorage = Store
	v.Cfg.QuanFilesInDirectory = len(file_handler.GetFilesInDirectory(v.Cfg.PathToMethodsDirectory))

	Chad := change_detector.NewChad(&v.Cfg)
	Fired := file_readiness_detector.NewFired(&v.Cfg)
	Infinit := injection_of_function_init.NewInfinit(&v.Cfg)

	mechanismsStart(channel, Chad, Fired, Infinit)
	for {
		select {
		case val := <-channel.ContentChange:
			{
				if val {
					fmt.Println("channel is catch the signal")
				}
			}
		case err := <-channel.Errors:
			{
				if err != nil {
					log.Println(err)
				}
			}
		default:

		}
	}
}

func mechanismsStart(channel *models.Channel, mechanisms ...Mechanisms) {
	for _, el := range mechanisms {
		el.Launch(*channel)
	}
}
