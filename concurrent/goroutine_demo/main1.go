package main

import (
	"fmt"
	"sync"
)

var group = sync.WaitGroup{}

func say(i int) {
	group.Add(1)
	fmt.Printf("%d -------娜扎。。。\n", i)
	group.Done()
}
func main() {

	for i := 0; i < 1000; i++ {
		go say(i)
	}
	fmt.Println("-------main。。。")
	group.Wait()

}
