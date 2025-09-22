package services

import (
	"context"
	"fmt"
	"time"
)

// 設定一個 2 秒後自動取消的 context
func Context1() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan bool)

	go func() {
		// 模擬耗時任務
		time.Sleep(3 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("任務完成 ✅")
	case <-ctx.Done():
		fmt.Println("任務被取消 ❌:", ctx.Err())
	}
}

// 建議自定義 key 避免與其他套件衝突（避免直接用 string）
type contextKey string

const (
	UserIDKey    contextKey = "userID"
	RequestIDKey contextKey = "requestID"
)

func handleRequest(ctx context.Context) {
	// 取得數據
	userID := ctx.Value(UserIDKey).(int)
	requestID := ctx.Value(RequestIDKey).(string)

	fmt.Println("處理請求中...")
	fmt.Println("使用者 ID:", userID)
	fmt.Println("請求 ID:", requestID)

	// 模擬呼叫下游服務
	processData(ctx)
}

func processData(ctx context.Context) {
	// goroutine 也能取得 context 中的資料
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("子 goroutine 讀取 userID:", ctx.Value(UserIDKey))
		fmt.Println("子 goroutine 讀取 requestID:", ctx.Value(RequestIDKey))
	}()

	time.Sleep(1 * time.Second) // 等 goroutine 結束
}

// 建立一個背景 context，並塞入數據
func Context2() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIDKey, 42)
	ctx = context.WithValue(ctx, RequestIDKey, "REQ-123456")

	// 將 context 傳遞給處理請求的函式
	handleRequest(ctx)
}
