package main

import (
	"backend/handlers"
	"backend/middleware"
	"backend/redis"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()

	rdb, err := redis.RedisInit()
	if err != nil {
		fmt.Println(err)
	}

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Overrides other settings
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false, // Disabled to avoid wildcard + credentials conflict
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.RedisPool(rdb))

	r.POST("/shorten", handlers.GetUrl)
	r.GET("/redirect/:id", handlers.RedirectUser)

	r.Run(":8080")

}
