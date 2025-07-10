package logger

import (
	"concurrency_task/internal/channels"
	"concurrency_task/internal/config"
	"fmt"
)

type Logger struct {
}

func NewLogger(cfg config.Config) (*Logger, error) {
	return &Logger{}, nil
}

func (l *Logger) Launch(channel channels.Channel) {
	go func() {
		for {
			select {
			case err := <-channel.ReadErrorsFromChannel():
				fmt.Println(err)
			}
		}
	}()
}
