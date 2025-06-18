package capabilityChecker

import (
	"fmt"
	"os"
	"time"
)

type CapChecker struct {
	PathToMethodsDirectory string
	Interval               time.Duration
	QuantityMethods        int
}

func (cc *CapChecker) LaunchChecker(channel chan bool) {
	ticker := time.NewTicker(cc.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if cc.QuantityMethods != cc.checkNewFunctions() {
				l, _ := os.ReadDir(cc.PathToMethodsDirectory)
				cc.QuantityMethods = len(l)
				channel <- true
			}
		}
	}
}

func (cc *CapChecker) checkNewFunctions() int {
	dirElements, err := os.ReadDir(cc.PathToMethodsDirectory)
	if err != nil {
		fmt.Println(err)
	}
	return len(dirElements)
}
