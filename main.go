package main

import (
	"fmt"
	"math/rand"
	"time"
)

// @title My Gin API
// @version 1.0
// @description Example API using Gin, JWT, GORM, Swagger

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// func main() {
// 	core.InitDB()
// 	core.DB.AutoMigrate(&models.User{})
// 	r := routes.SetupRouter()
// 	r.Run(":3000")
// }

// 程式碼修改自 Concurrency Patterns in Go: sync.WaitGroup @ https://www.calhoun.io/

func main() {
	notify("Service-1", "Service-2", "Service-3")
}

func notifying(res chan string, s string) {
	fmt.Printf("Starting to notifying %s...\n", s)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	res <- fmt.Sprintf("Finish notifying %s", s)
}

func notify(services ...string) {
	res := make(chan string)
	var count int = 0

	for _, service := range services {
		count++
		go notifying(res, service)
	}

	for i := 0; i < count; i++ {
		fmt.Println(<-res)
	}

	fmt.Println("All service notified!")
}
