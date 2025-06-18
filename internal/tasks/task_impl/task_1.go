package task_impl

import (
	"fmt"
	"sync"
)

type Task_1 struct {
	Number int
}

func NewTask_1() Task_1 {
	return Task_1{
		Number: 1,
	}
}
func (t Task_1) Launch() {
	counter := 20
	wg := sync.WaitGroup{}
	for i := 0; i <= counter; i++ {
		//i := i Shadowing
		wg.Add(1)
		go func(in int) {
			defer wg.Done()
			fmt.Println(in * in)
		}(i)
	}

	wg.Wait()
}

func init() {
	
}
