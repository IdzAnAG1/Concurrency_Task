package errors_handler

import (
	"concurrency_task/internal/channels"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type ErrorsHandler struct {
	logger zap.Logger
}

func NewErrorsHandler(logger zap.Logger) *ErrorsHandler {
	return &ErrorsHandler{
		logger: logger,
	}
}
func (eh *ErrorsHandler) Launch(ctx context.Context, channel channels.Channel) {
	go func() {
		eh.logger.Info("Errors Handler was launched")
		for {
			select {
			case err := <-channel.ReadErrorsFromChannel():
				eh.logger.Error(err.Error())
			case <-ctx.Done():
				eh.logger.Info("The completion signal is received in Errors Handler")
				return
			}
		}
	}()
}
