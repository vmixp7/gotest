package services

import (
	"fmt"
	"sync"
	"time"
)

func printHello() {
	fmt.Println("Hello World")
}

func Gotest1() {
	fmt.Println("main execution started")

	// call function
	go printHello()

	// block here
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main execution stopped")
}

func Gotest2() {
	fmt.Println("main() started")

	c := make(chan string)

	// anonymous goroutine
	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "John"
	fmt.Println("main() ended")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine 完成时调用 Done()
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d finished\n", id)
}

func Gotest3() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数器
		go worker(i, &wg)
	}

	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("All workers done")
}
