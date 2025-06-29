package main

import (
	_ "concurrency_task/internal/tasks/task_impl"
	"concurrency_task/internal/utils/general"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	userType := ""
	re := regexp.MustCompile(`(?)type .*\b struct\b`)
	re1 := regexp.MustCompile(`(?)func init\(\) \{`)
	files := general.GetFilesInDirectory("internal/tasks/task_impl")
	str := strings.Split(general.ReadFromFile("internal/tasks/task_impl", files[0]), "\n")

	for i, v := range str {
		if re.MatchString(v) {
			fmt.Println(i)
		}
	}
	for i, v := range str {
		if re1.MatchString(v) {
			fmt.Println(i)
		}
	}
	fmt.Println(userType)
	/*	var (
			BooleanChannel = make(chan bool)
		)
		Tasks := task_storage.GetStorageInstance()
		go func(store task_storage.TaskStorage, booleanChannel chan bool) {
			Chad := chad.NewChad("internal/tasks/task_impl", 500*time.Millisecond, len(general.GetFilesInDirectory("internal/tasks/task_impl")))
			fmt.Println("Cap Checker Started")
			Chad.Launch(booleanChannel)
		}(*Tasks, BooleanChannel)
		for {
			select {
			case val := <-BooleanChannel:
				{
					if val {
						fmt.Println("channel is catch the signal")
					}
				}
			default:

			}
		}*/
}
