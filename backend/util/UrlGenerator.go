package util

import (
	"backend/model"
	"context"
	"crypto/rand"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"math/big"
	"strconv"
	"time"
)

var ctx = context.Background()

func GenerateURL(c *gin.Context, url model.Url) string {

	log.Println("inside GenerateURL")

	val, err := countId(c, url)
	if err != nil {
		return ""
	}

	shortCode := val

	shortUrl := "http://localhost:8080/redirect/" + shortCode

	return shortUrl

}

func countId(c *gin.Context, url model.Url) (string, error) {

	log.Println("Inside countId")

	redisInterface, exists := c.Get("client")
	if !exists {
		log.Println("redisInterface does not exist")
		return "", errors.New("redisInterface does not exist")
	}

	redisClient := redisInterface.(*redis.Client)
	urlCode, _ := generateRandomURL()

	err := redisClient.Set(ctx, urlCode, url.OriginalUrl, 10*time.Minute).Err()
	if err != nil {
		log.Println(err)
		return "", errors.New("count does not exist")
	}

	return urlCode, nil

}

func generateRandomURL() (string, error) {
	max := big.NewInt(1e10)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(n.Int64(), 36), nil
}
