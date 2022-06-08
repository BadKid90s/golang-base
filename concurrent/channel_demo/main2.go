package main

import (
	"fmt"
	"sync"
)

func write2(ch1 chan int) {
	for i := 0; i < 2000; i++ {
		ch1 <- i
	}
	close(ch1)
	group2.Done()
}

func readValue(ch1 chan int, reschan chan int) {
	for i := range ch1 {
		sum := sum(i)
		reschan <- sum
	}
	close(reschan)
	group2.Done()
}

func sum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
var group2=sync.WaitGroup{}

func main2() {
	intChan := make(chan int, 2000)
	resChan := make(chan int, 2000)

	group2.Add(2)

	go write2(intChan)

	go readValue(intChan, resChan)

	for i := range resChan {
		fmt.Println(i)
	}
	group2.Wait()
}
