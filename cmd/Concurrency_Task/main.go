package main

import (
	"concurrency_task/internal/models"
	_ "concurrency_task/internal/tasks/task_impl"
	"concurrency_task/internal/utils/chad"
	"concurrency_task/internal/utils/fired"
)

func main() {
	// Запуск механизмов и окружения(в виде структуры каналов) для них
	channels := models.NewChannel()
	Fired := fired.NewFired()
	Chad := chad.NewChad()

	Fired.Launch()
	Chad.Launch()
	//contentChannel := make(chan string)
	//f := fired.NewFired()
	//
	//f.Run(contentChannel)
	//files := general.GetFilesInDirectory("internal/tasks/task_impl")
	//str := general.ReadFromFile("internal/tasks/task_impl", files[0])
	//contentChannel <- str
	//
	//var (
	//	BooleanChannel = make(chan bool)
	//)
	//
	//Tasks := task_storage.GetStorageInstance()
	//
	//go func(store task_storage.TaskStorage, booleanChannel chan bool) {
	//	Chad := chad.NewChad("internal/tasks/task_impl", 500*time.Millisecond, len(general.GetFilesInDirectory("internal/tasks/task_impl")))
	//	fmt.Println("Cap Checker Started")
	//	Chad.Launch(booleanChannel)
	//}(*Tasks, BooleanChannel)
	//for {
	//	select {
	//	case val := <-BooleanChannel:
	//		{
	//			if val {
	//				fmt.Println("channel is catch the signal")
	//			}
	//		}
	//	default:
	//
	//	}
	//}
}
