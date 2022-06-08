package main

import (
	"fmt"
	"runtime"
	"sync"
)

func a() {
	for i := 0; i < 100; i++ {
		fmt.Println("a------------", i)
	}
	group.Done()
}
func b() {
	for i := 0; i < 100; i++ {
		fmt.Println("b------------", i)
	}
	group.Done()
}

var group = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(2)
	group.Add(2)
	go a()
	go b()
	group.Wait()
}
