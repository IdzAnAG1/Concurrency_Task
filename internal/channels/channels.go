package channels

import "concurrency_task/internal/models"

type Channel struct {
	content          chan string
	contentChange    chan bool
	contentIndicator chan *models.InfinitData
	errors           chan error
}

func NewChannel() *Channel {
	return &Channel{
		content:          make(chan string),
		contentChange:    make(chan bool),
		contentIndicator: make(chan *models.InfinitData),
		errors:           make(chan error),
	}
}

func (c *Channel) SendToContentChannel(element string) {
	c.content <- element
}
func (c *Channel) SendToChangeChannel(element bool) {
	c.contentChange <- element
}
func (c *Channel) SendToChannelContentIndicator(element *models.InfinitData) {
	c.contentIndicator <- element
}
func (c *Channel) SendErrorsToChannel(err error) {
	c.errors <- err
}
func (c *Channel) ReadContentFromChannel() <-chan string {
	return c.content
}
func (c *Channel) ReadErrorsFromChannel() <-chan error {
	return c.errors
}
func (c *Channel) ReadInfDataFromChannel() <-chan *models.InfinitData {
	return c.contentIndicator
}
func (c *Channel) ReadChangeFromChannel() <-chan bool {
	return c.contentChange
}
