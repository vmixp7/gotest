package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

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

func GetLock(c *gin.Context) {
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

	c.JSON(http.StatusOK, released)
}
