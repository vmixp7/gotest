package main

import (
	"fmt"
	"gotest/libary"
)

// @title My Gin API
// @version 1.0
// @description Example API using Gin, JWT, GORM, Swagger

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

//	func main() {
//		core.InitDB()
//		core.DB.AutoMigrate(&models.User{})
//		r := routes.SetupRouter()
//		r.Run(":3000")
//	}

func squares(c chan int) {
	// 把 0 ~ 9 寫入 channel 後便把 channel 關閉
	for i := 0; i <= 8; i++ {
		c <- i
	}

	close(c)
}

func main() {
	fmt.Println("main() started")
	// c := make(chan int)

	// // 發動 squares goroutine
	// go squares(c)

	// // 監聽 channel 的值：週期性的 block/unblock main goroutine 直到 squares goroutine close
	// for val := range c {
	// 	fmt.Println(val)
	// }

	makeCounter := libary.MakeCounter()
	for i := 0; i < 10; i++ {
		fmt.Println(makeCounter())
	}

	fmt.Println("main() close")
}
