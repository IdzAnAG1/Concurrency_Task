package interruptor

import (
	"concurrency_task/internal/channels"
	"context"
	"os"
	"os/signal"
	"syscall"
)

type Interruptor struct {
	chans  channels.Channel
	cancel context.CancelFunc
}

func NewInterruptor(channel channels.Channel, cancelFunc context.CancelFunc) *Interruptor {
	return &Interruptor{
		chans:  channel,
		cancel: cancelFunc,
	}
}
func (i *Interruptor) Run() {
	go func() {
		signal.Notify(i.chans.GetInterruptionChannel(), os.Interrupt, syscall.SIGTERM)
		<-i.chans.GetInterruptionChannel()
		i.cancel()
	}()
}
