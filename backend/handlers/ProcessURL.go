package handlers

import (
	"backend/model"
	"backend/util"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

func GetUrl(c *gin.Context) {

	var url model.Url

	log.Println("inside GetUrl")

	if err := c.ShouldBindJSON(&url); err != nil {
		log.Println(err)
	}
	log.Println(url)
	url.ShortUrl = ProcessURL(c, url)

	log.Println("url:", url.ShortUrl)

	c.JSON(http.StatusOK, gin.H{"url": url.ShortUrl})

}

func ProcessURL(c *gin.Context, url model.Url) string {

	log.Println("inside ProcessURL", "URL: ", url.OriginalUrl)

	redisInterface, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "redis client not found"})
		return ""
	}

	client := redisInterface.(*redis.Client)

	url.ShortUrl = util.GenerateURL(c, url)

	err := client.Set(c.Request.Context(), url.OriginalUrl, url.ShortUrl, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return ""
	}

	return url.ShortUrl

}
