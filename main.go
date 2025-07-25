package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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

// https://pkg.go.dev/github.com/gin-contrib/sessions?tab=overview#cookie-based
// cookie-based
func main() {
	r := gin.Default()
	// 建立 store
	store := cookie.NewStore([]byte("secret"))

	// session 的名稱會在 browser 變成 cookie 的 key
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		// 從 ctx 中取出 session
		session := sessions.Default(c)
		var count int

		// 取得 session 中的值
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}

		// 設定 session 中的值（不會儲存）
		session.Set("count", count)

		// 儲存 session 中的值
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
