package services

import (
	"fmt"
	"time"
)

func squares(c chan int) {
	// 把 0 ~ 9 寫入 channel 後便把 channel 關閉
	for i := 0; i <= 8; i++ {
		c <- i
	}
	close(c) //沒有關閉的話會deadlock
}

func ChannelTest1() {
	c := make(chan int)

	// 發動 squares goroutine
	go squares(c)

	// 監聽 channel 的值：週期性的 block/unblock main goroutine 直到 squares goroutine close
	for val := range c {
		fmt.Println(val)
	}
}

// 隨機選取
func ChannelTes2() {
	ch := make(chan int, 1) //buffered channel會讀取所有的
	ch <- 1
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	case <-ch:
		fmt.Println("random 03")
	}
}

func ChannelTest3() {
	ch := make(chan int)

	select {
	case <-ch:
		fmt.Println("receive value from channel")

	//	超過一秒沒有收到主要 channel 的 value，就會收到 time.After 送來的訊息
	case <-time.After(1 * time.Second):
		fmt.Println("timeout after 1 second")
	}
}

func service1() {
	fmt.Println("Hello from service1 ")
	time.Sleep(2 * time.Second)
}
func service2() {
	fmt.Println("Hello from service2 ")
	time.Sleep(5 * time.Second)
}

func ChannelTest4() {
	fmt.Println("main() started")

	go service1()
	go service2()

	// 這個 select 會永遠 block 在這，service1 和 service2 輪流輸出訊息
	select {}

	fmt.Println("main() stopped")
}

func ChannelTest5() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return // 如果沒有 return 的話程式將不會結束，一直卡在 for loop 中
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
