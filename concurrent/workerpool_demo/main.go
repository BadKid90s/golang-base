package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker Id:%d, start job:%d \n", id, job)
		results <- job * job
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 100)

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < 5; i++ {
		num := <-results
		fmt.Println("", num)
	}
}
