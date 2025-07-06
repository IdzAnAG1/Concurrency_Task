package models

import "sync"

type Channel struct {
	Content          chan string
	ContentChange    chan bool
	ContentIndicator chan *InfinitData_v2
	Errors           chan error
}

func NewChannel() *Channel {
	return &Channel{
		Content:          make(chan string),
		ContentChange:    make(chan bool),
		ContentIndicator: make(chan *InfinitData_v2),
	}
}

type ReadinessIndicator struct {
	UserStructIsExist              bool
	InterfaceImplementationIsExist bool
	InitFuncIsExist                bool
}
type ReadinessIndicator_v2 struct {
	mu           sync.Mutex
	FileFullness map[string]int
}

func NewReadinessIndicator() *ReadinessIndicator_v2 {
	return &ReadinessIndicator_v2{
		FileFullness: make(map[string]int),
	}
}
func (r *ReadinessIndicator_v2) Put(exp string, index int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.FileFullness[exp] = index
}

type InfinitData struct {
	FileName  string
	Indicator ReadinessIndicator
}
type InfinitData_v2 struct {
	FileName  string
	Indicator ReadinessIndicator_v2
}
