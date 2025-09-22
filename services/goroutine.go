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

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine 完成时调用 Done()
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d finished\n", id)
}

func Gotest3() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数器
		go work(i, &wg)
	}

	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("All workers done")
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // 模擬工作
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

// 工作池範例,限制同時工作的 goroutine 數量
func WorkerPull() {
	const numJobs = 10   // 總共的工作數量
	const numWorkers = 5 // 限制同時工作的 goroutine 數量

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// 啟動 workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results, &wg)
	}

	// 發送 job
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 等待所有 workers 完成
	wg.Wait()
	close(results)

	// 處理結果
	for r := range results {
		fmt.Println("Result:", r)
	}
}
