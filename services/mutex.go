package services

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
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

// 1. acquire_timeout：每個人能獲得鎖的上限時間為30s
// 2. lock_timeout：鎖的過期時間為10s
var ctx = context.Background()

func acquireLock(rdb *redis.Client, key string, value string, expiration time.Duration) (bool, error) {
	// 將 NX（不存在才設值）和 EX（過期時間）加上，避免死鎖
	result, err := rdb.SetNX(ctx, key, value, expiration).Result()
	return result, err
}

func releaseLock(rdb *redis.Client, key string, value string) (bool, error) {
	// 透過 Lua 腳本確保只有持有者才能刪除鎖
	luaScript := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
    `
	result, err := rdb.Eval(ctx, luaScript, []string{key}, value).Result()
	return result.(int64) == 1, err
}

// 分散式鎖, 以 Redis 為例, 搶票系統
func Mutex5() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	lockKey := "my-lock"
	lockValue := "unique-request-id-123"
	lockTTL := 10 * time.Second

	ok, err := acquireLock(rdb, lockKey, lockValue, lockTTL)
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("Failed to acquire lock")
		return
	}

	fmt.Println("Lock acquired, doing critical work...")
	time.Sleep(3 * time.Second)

	released, err := releaseLock(rdb, lockKey, lockValue)
	if err != nil {
		panic(err)
	}

	if released {
		fmt.Println("Lock released successfully")
	} else {
		fmt.Println("Failed to release lock")
	}
}

// sync.Map 的使用
// sync.Map 是 Go 語言標準庫中提供的一種並發安全的映射（map）結構，適用於高併發場景下的鍵值對存取。
// 它與傳統的 map 不同，因為它內部已經實現了鎖機制，讓多個 goroutine 可以安全地同時讀寫資料，而不需要額外的鎖。
func SyncMap() {
	var m sync.Map

	// 1. Store: 儲存鍵值對
	fmt.Println("Storing key-value pairs...")
	m.Store("name", "Alice")
	m.Store("age", 30)

	// 2. Load: 載入鍵值
	// Load方法會回傳兩個值：第一個是值，第二個是布林值，表示該鍵是否存在
	if name, ok := m.Load("name"); ok {
		fmt.Printf("Loaded 'name': %s\n", name)
	}

	// 3. LoadOrStore: 載入或儲存
	// 如果鍵存在，則返回現有的值；如果不存在，則儲存並返回新值
	fmt.Println("\nUsing LoadOrStore...")
	// 鍵 "country" 不存在，所以會儲存 "Taiwan"
	if value, loaded := m.LoadOrStore("country", "Taiwan"); !loaded {
		fmt.Printf("Key 'country' didn't exist, stored '%s'.\n", value)
	}

	// 鍵 "name" 已經存在，會返回舊值 "Alice"，不會儲存 "Bob"
	if value, loaded := m.LoadOrStore("name", "Bob"); loaded {
		fmt.Printf("Key 'name' already existed, loaded '%s'.\n", value)
	}

	// 4. Delete: 刪除鍵
	fmt.Println("\nDeleting key 'age'...")
	m.Delete("age")
	if _, ok := m.Load("age"); !ok {
		fmt.Println("Key 'age' has been deleted.")
	}

	// 5. Range: 遍歷所有鍵值對
	// Range 方法需要一個回呼函式，如果回呼函式返回 false，則會停止遍歷
	fmt.Println("\nIterating over the map:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		// 如果你想停止遍歷，可以在這裡回傳 false
		return true
	})
}
