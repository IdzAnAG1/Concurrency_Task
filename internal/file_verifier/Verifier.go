package file_verifier

import (
	"concurrency_task/internal/channels"
	"concurrency_task/internal/config"
	"concurrency_task/internal/file_verifier/change_detector"
	"concurrency_task/internal/file_verifier/errors_handler"
	"concurrency_task/internal/file_verifier/file_readiness_detector"
	"concurrency_task/internal/file_verifier/injection_of_function_init"
	"concurrency_task/internal/interruptor"
	"concurrency_task/internal/logger"
	"concurrency_task/internal/tasks/task_code_storage"
	_ "concurrency_task/internal/tasks/task_impl"
	"concurrency_task/internal/utils/file_handler"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"os"
	"sync"
	"time"
)

type MechanismsStarter interface {
	Launch(context.Context, *sync.WaitGroup)
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

	channel := *channels.NewChannel(*logs)
	ctx, cancel := context.WithCancel(context.Background())
	interuptor := interruptor.NewInterruptor(channel, cancel)
	interuptor.Run()

	v.TCStorage = v.StoreInitializer(logs, channel)
	v.QuanFilesInDirectory = len(v.QuantityFilesUpdater(channel))

	wg := sync.WaitGroup{}

	Chad := v.ChadInitializer(logs, channel)
	Fired := v.FiredInitializer(logs, channel)
	Infinit := v.InfinitInitializer(logs, channel)
	errHan := v.ErrHanInitializer(logs, channel)

	mechanismsStart(ctx, &wg, errHan, Chad, Fired, Infinit)

loop:
	for {
		select {
		case <-ctx.Done():
			logs.Info("The completion signal has been received")
			break loop
		}
	}
	wg.Wait()

	channel.CloseChannels()
	return nil
}
func (v *Verifier) QuantityFilesUpdater(channel channels.Channel) []os.DirEntry {
	files, err := file_handler.GetFilesInDirectory(v.PathToMethodsDirectory)
	if err != nil {
		channel.SendErrorsToChannel(err)
	}
	return files
}
func (v *Verifier) StoreInitializer(logger *zap.Logger, channel channels.Channel) *task_code_storage.TCStorage {
	Store := task_code_storage.NewTCStorage(*logger)
	err := Store.Initialize(v.PathToMethodsDirectory)
	if err != nil {
		channel.SendErrorsToChannel(err)
	}
	return Store
}
func (v *Verifier) ErrHanInitializer(logger *zap.Logger, channel channels.Channel) *errors_handler.ErrorsHandler {
	return errors_handler.NewErrorsHandler(
		*logger,
		channel,
	)
}
func (v *Verifier) InfinitInitializer(logger *zap.Logger, channel channels.Channel) *injection_of_function_init.Infinit {
	return injection_of_function_init.NewInfinit(
		*logger,
		channel,
		v.PathToMethodsDirectory,
		v.TCStorage,
	)
}
func (v *Verifier) ChadInitializer(logger *zap.Logger, channel channels.Channel) *change_detector.ChaD {
	return change_detector.NewChad(
		*logger,
		channel,
		v.PathToMethodsDirectory,
		v.Interval,
		v.QuanFilesInDirectory,
		v.TCStorage,
	)
}
func (v *Verifier) FiredInitializer(logger *zap.Logger, channel channels.Channel) *file_readiness_detector.Fired {
	return file_readiness_detector.NewFired(
		*logger,
		channel,
		v.TCStorage,
	)
}
func mechanismsStart(ctx context.Context, group *sync.WaitGroup, mechanisms ...MechanismsStarter) {
	for _, el := range mechanisms {
		el.Launch(ctx, group)
	}
}
