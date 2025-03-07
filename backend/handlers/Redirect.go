package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

func RedirectUser(c *gin.Context) {
	param := c.Param("id")

	redisInterface, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Redis client not found"})
		return
	}

	client := redisInterface.(*redis.Client)

	url, err := client.Get(c.Request.Context(), param).Result()
	if err == redis.Nil || url == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Short URL not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Redis error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
