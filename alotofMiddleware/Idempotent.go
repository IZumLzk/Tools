package alotofMiddleware

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParseToken(v any) (map[string]interface{}, error) {
	return map[string]interface{}{}, errors.New("aaa")
}
func UserlistMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "请先登录",
				"code": 400,
			})
			c.Abort()
			return
		}
		parseToken, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "token解析失败",
				"code": 400,
			})
			c.Abort()
			return
		}
		uid := parseToken["userId"].(string)
		//redisKey := "user:list"
		////ok, _ := redis2.RC.SetNX(context.Background(), redisKey, "1", 5*time.Second)//.Result()
		////if !ok {
		////	c.JSON(http.StatusBadRequest, gin.H{
		////		"code": 400,
		////		"msg":  "error",
		////	})
		////	c.Abort()
		////	return
		////}

		c.Set("userId", uid)
		c.Next()
	}
}
func CacheHeader() func(c *gin.Context) {
	return func(c *gin.Context) {
		contentVersion := "v0.0.1"
		etag := fmt.Sprintf("/m%x", md5.Sum([]byte(contentVersion)))

		ClientEtag := c.GetHeader("If-None-Match")
		if ClientEtag == etag {
			c.AbortWithStatus(http.StatusNotModified)
			return
		}
		c.Header("Cache-Control", "no-cache")
		c.Header("Etag", etag)
		c.Next()
	}
}
