package main

import "fmt"

// 判断是否是素数
func isPrimeNum(n int) bool {
	if n <= 3 {
		return n > 1
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//素数写入管道
func primeNum(intChan, primeChan chan int, exitChan chan bool) {
	for {
		x, ok := <-intChan
		if !ok {
			break
		}
		if isPrimeNum(x) {
			primeChan <- x
		}
	}
	fmt.Println("有一个primeNum协程因为取不到数据，退出。")
	exitChan <- true
}
func main3() {
	//原数字
	intChan := make(chan int, 1000)
	//素数管道
	primeChan := make(chan int, 2000)
	//退出管道
	exitChan := make(chan bool, 4)

	//开启一个协程写入数字
	go func(intChan chan int) {
		for i := 1; i <= 8000; i++ {
			intChan <- i
		}
		close(intChan)
	}(intChan)

	//开启4个协程，从管道取数，
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
		close(exitChan)
	}()

	for i := range primeChan {
		fmt.Println("-------", i)
	}

	//var num = 0
	//for b := range exitChan {
	//	if b {
	//		num++
	//	}
	//	if num == 4 {
	//		close(primeChan)
	//		close(exitChan)
	//		break
	//	}
	//}
}
