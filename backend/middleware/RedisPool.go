package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RedisPool(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("client", client)
		c.Next()
	}
}
