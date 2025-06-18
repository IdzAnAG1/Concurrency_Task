package main

import (
	"concurrency_task/internal/utils/capabilityChecker"
	"fmt"
	"os"
	"time"
)

func main() {

}

func temp() {
	channel := make(chan bool)
	l, _ := os.ReadDir("internal/tasks/task_impl")
	CC := capabilityChecker.CapChecker{
		PathToMethodsDirectory: "internal/tasks/task_impl",
		Interval:               3 * time.Second,
		QuantityMethods:        len(l),
	}
	go CC.LaunchChecker(channel)

	for {
		select {
		case val := <-channel:
			{
				if val {
					fmt.Println("New task")
					return
				}
			}
		default:
			fmt.Println("Ничего нового")
		}
	}
}
