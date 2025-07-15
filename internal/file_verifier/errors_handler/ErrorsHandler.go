package errors_handler

import (
	"concurrency_task/internal/channels"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"sync"
)

type ErrorsHandler struct {
	channels channels.Channel
	logger   zap.Logger
}

func NewErrorsHandler(logger zap.Logger, channels channels.Channel) *ErrorsHandler {
	return &ErrorsHandler{
		channels: channels,
		logger:   logger,
	}
}
func (eh *ErrorsHandler) Launch(ctx context.Context, group *sync.WaitGroup) {
	group.Add(1)
	go func() {
		defer group.Done()
		eh.logger.Info("Errors Handler was launched")
		for {
			select {
			case err := <-eh.channels.ReadErrorsFromChannel():
				eh.logger.Error(err.Error())
			case <-ctx.Done():
				eh.logger.Info("The completion signal is received in Errors Handler")
				return
			}
		}
	}()
}
