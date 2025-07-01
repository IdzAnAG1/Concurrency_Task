package models

type Channel struct {
	Content chan string
	Boolean chan bool
}

func NewChannel() *Channel {
	return &Channel{
		Content: make(chan string),
		Boolean: make(chan bool),
	}
}
