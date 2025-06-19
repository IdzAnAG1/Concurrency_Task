package capabilityChecker

import (
	"concurrency_task/internal/utils/general"
	"fmt"
	"os"
	"sync"
	"time"
)

type CapChecker struct {
	mu                     sync.Mutex
	PathToMethodsDirectory string
	Interval               time.Duration
	FunctionsMap           map[string]string
}

func (cc *CapChecker) LaunchChecker(channel chan bool) {
	ticker := time.NewTicker(cc.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if cc.isMapWasUpdated() {
				channel <- true
			}
		}
	}
}

func (cc *CapChecker) isMapWasUpdated() bool {
	FITDir, err := os.ReadDir(cc.PathToMethodsDirectory) // FITDir - Files In The Directory
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range FITDir {
		cc.mu.Lock()

		content, errCont := general.GetFileContents(cc.PathToMethodsDirectory, f.Name()) // Получение содержания текущего файла
		if errCont != nil {
			fmt.Println(errCont)
		}
		hash := general.ConvertToHash(content) // Перевод вв Хэш содержимое

		if _, exists := cc.FunctionsMap[f.Name()]; !exists {
			cc.FunctionsMap[f.Name()] = hash
			return true // в случае если не было изначально подобнго файла ,то добавляем его в мапу, и говорим о том что мапа была изменена
		} else {
			if cc.FunctionsMap[f.Name()] != hash {
				cc.FunctionsMap[f.Name()] = hash
				return true // В случае если подобный файл бы в мапе то сверяем его хэш, если хэш изменен то присваиваем новый хэщ, если нет то просто пропускаем ход
			}
		}
		cc.mu.Unlock()
	}
	return false
}
