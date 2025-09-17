package services

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func GcTest() {
	// 調整這裡來比較不同的 GOGC 效果，例如 50 / 100 / 200 / off
	debug.SetGCPercent(100) // 預設是 100

	fmt.Println("Start GC Test with GOGC =", debug.SetGCPercent(-1))
	debug.SetGCPercent(100) // 再設定回來

	// 每輪會產生一堆垃圾
	for i := 0; i < 5; i++ {
		workload()
		printMemStats()
		time.Sleep(2 * time.Second)
	}
}

func workload() {
	// 建立大量短命物件 (垃圾)
	for i := 0; i < 10000000; i++ {
		_ = make([]byte, 1024) // 每個 1KB，10 萬個就是 100MB
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %v MB, NumGC = %v\n", m.HeapAlloc/1024/1024, m.NumGC)
}
