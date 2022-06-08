package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Println("接受数据", x)
		case ch1 <- i:
			fmt.Println("发送数据", i)
		default:
			fmt.Println("-------------")
		}
	}

}
