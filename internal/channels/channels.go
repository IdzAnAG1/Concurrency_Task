package channels

import (
	"concurrency_task/internal/models"
	"go.uber.org/zap"
	"os"
)

type Channel struct {
	logger           zap.Logger
	content          chan string
	contentChange    chan bool
	contentIndicator chan *models.InfinitData
	errors           chan error
	interruption     chan os.Signal
}

func NewChannel(logger zap.Logger) *Channel {
	logger.Info("Channels is opened")
	return &Channel{
		logger:           logger,
		content:          make(chan string),
		contentChange:    make(chan bool),
		contentIndicator: make(chan *models.InfinitData),
		errors:           make(chan error),
		interruption:     make(chan os.Signal, 1),
	}
}
func (c *Channel) CloseChannels() {
	go func() {
		close(c.content)
		close(c.contentChange)
		close(c.contentIndicator)
		close(c.errors)
	}()
	c.logger.Info("Channels Closed")
}
func (c *Channel) GetInterruptionChannel() chan os.Signal {
	return c.interruption
}
func (c *Channel) SendToContentChannel(element string) {
	c.logger.Info("Data has been sent to the content channel")
	c.content <- element
}
func (c *Channel) SendToChangeChannel(element bool) {
	c.logger.Info("Data has been sent to the Change channel")
	c.contentChange <- element
}
func (c *Channel) SendToChannelContentIndicator(element *models.InfinitData) {
	c.logger.Info("Data has been sent to the ContentIndicator channel")
	c.contentIndicator <- element
}
func (c *Channel) SendErrorsToChannel(err error) {
	c.errors <- err
}
func (c *Channel) ReadContentFromChannel() <-chan string {
	c.logger.Info("Data is read from the Content channel")
	return c.content
}
func (c *Channel) ReadErrorsFromChannel() <-chan error {
	return c.errors
}
func (c *Channel) ReadInfDataFromChannel() <-chan *models.InfinitData {
	c.logger.Info("Data is read from the InfinitData channel")
	return c.contentIndicator
}
func (c *Channel) ReadChangeFromChannel() <-chan bool {
	c.logger.Info("Data is read from the ContentChange channel")
	return c.contentChange
}
