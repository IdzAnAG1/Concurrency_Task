package models

import "sync"

type Channel struct {
	Content          chan string
	ContentChange    chan bool
	ContentIndicator chan *InfinitData
	Errors           chan error
}

func NewChannel() *Channel {
	return &Channel{
		Content:          make(chan string),
		ContentChange:    make(chan bool),
		ContentIndicator: make(chan *InfinitData),
	}
}

type ReadinessIndicator struct {
	mu           sync.Mutex
	FileFullness map[string]int
}

func NewReadinessIndicator() *ReadinessIndicator {
	return &ReadinessIndicator{
		mu:           sync.Mutex{},
		FileFullness: make(map[string]int),
	}
}
func (r *ReadinessIndicator) Put(exp string, index int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.FileFullness[exp] = index
}

type InfinitData struct {
	FileName  string
	Indicator ReadinessIndicator
}
