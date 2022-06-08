package main

import (
	"fmt"
	"sync"
)

var (
	x         int64
	waitGroup sync.WaitGroup
	lock      sync.Mutex
)

func mutexLock() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(2)

	go mutexLock()
	go mutexLock()

	waitGroup.Wait()

	fmt.Println(x)
}
