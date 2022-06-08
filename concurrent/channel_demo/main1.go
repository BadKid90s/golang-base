package main

import (
	"fmt"
	"sync"
)

func write(ch1 chan int) {
	for i := 0; i < 50; i++ {
		ch1 <- i
		fmt.Printf("写数据：%d \n", i)
	}
	close(ch1)
	group.Done()
}

func read(ch1 chan int) {
	// 在主goroutine中从ch2中接收值打印
	for i := range ch1 { // 通道关闭后会退出for range循环
		fmt.Printf("读数据：%d \n", i)
	}
	group.Done()
}

var group = sync.WaitGroup{}

func main1() {
	intChan := make(chan int, 50)

	group.Add(2)

	go write(intChan)

	go read(intChan)

	group.Wait()
}
