package main

import "fmt"
// 只发送的通道
func generateNum(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

// 只读取的通道，只发送的通道
func square(ch1 <-chan int, ch2 chan<- int) {
	for i := range ch1 {
		ch2 <- i * i
	}
	close(ch2)
}

func main() {
	intChan := make(chan int, 10)
	resChan := make(chan int, 100)

	go generateNum(intChan)

	go square(intChan, resChan)

	for i := range resChan {
		fmt.Println("计算的平方：","----------", i)
	}

}
