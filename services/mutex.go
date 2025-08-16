package services

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

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
