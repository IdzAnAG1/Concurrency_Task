package main

import (
	"concurrency_task/internal/utils/capabilityChecker"
	"concurrency_task/internal/utils/general"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	Map := initDir("internal/tasks/task_impl")
	channel := make(chan bool)
	CC := capabilityChecker.CapChecker{
		PathToMethodsDirectory: "internal/tasks/task_impl",
		Interval:               3 * time.Second,
		FunctionsMap:           Map,
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

func initDir(path string) (firstMap map[string]string) {
	firstMap = make(map[string]string)
	dirElements, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, dirEl := range dirElements {
		content, err := general.GetFileContents(path+"/", dirEl.Name())
		if err != nil {
			log.Fatal(err)
		}
		firstMap[dirEl.Name()] = general.ConvertToHash(content)
	}
	return
}
