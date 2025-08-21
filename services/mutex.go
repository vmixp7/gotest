package services

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 互斥鎖
func Mutex1() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final count:", count)
}

// 使用原子操作
func Mutex2() {

	var count int32
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&count, 1)
			val := atomic.LoadInt32(&count)
			fmt.Println("Current count:", val)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final count:", count)
}

var m *sync.RWMutex
var val = 0

func read(i int) {
	fmt.Println(i, "begin read")
	m.RLock()
	time.Sleep(1 * time.Second)
	fmt.Println(i, "val: ", val)
	time.Sleep(1 * time.Second)
	m.RUnlock()
	fmt.Println(i, "end read")
}

func write(i int) {
	fmt.Println(i, "begin write")
	m.Lock()
	val = 10
	fmt.Println(i, "val: ", val)
	time.Sleep(1 * time.Second)
	m.Unlock()
	fmt.Println(i, "end write")
}

func Mutex3() {
	m = new(sync.RWMutex)
	go read(1)
	go write(2)
	go read(3)
	time.Sleep(5 * time.Second)
}

func Mutex4() {
	runtime.GOMAXPROCS(1) // 限制 Go runtime 只使用 1 個 OS 執行緒執行 goroutine
	wg := sync.WaitGroup{}
	count := 10000
	wg.Add(count)
	ans := 0
	for i := 0; i < count; i++ {
		go func() {
			ans++ // (問題點)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ans)
}
