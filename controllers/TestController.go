package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	ch := make(chan int)

	// 啟動一個Goroutine，將數字從1到10傳遞到通道中

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 在主Goroutine中讀取通道中的數字，並將它們加總

	sum := 0
	for num := range ch {
		sum += num
	}
	fmt.Println("The sum is:", sum)

	c.JSON(http.StatusOK, sum)
}
