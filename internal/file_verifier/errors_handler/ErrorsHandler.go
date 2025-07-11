package errors_handler

import (
	"concurrency_task/internal/channels"
	"go.uber.org/zap"
)

type ErrorsHandler struct {
	logger zap.Logger
}

func NewErrorsHandler(logger zap.Logger) *ErrorsHandler {
	return &ErrorsHandler{
		logger: logger,
	}
}
func (eh *ErrorsHandler) Launch(channel channels.Channel) {
	go func() {
		for {
			select {
			case err := <-channel.ReadErrorsFromChannel():
				eh.logger.Error(err.Error())
			}
		}
	}()
}
