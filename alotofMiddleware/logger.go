package alotofMiddleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "" {
			parseToken, err := ParseToken(token)
			if err != nil {
				log.Println(err.Error())
			}
			uid := parseToken["userId"]

			c.Next()
			startTime := time.Now()
			duration := time.Since(startTime).Milliseconds()
			event := Event{
				UserId:   uid,
				Method:   c.Request.Method,
				Path:     c.Request.URL.Path,
				Duration: duration,
				Status:   c.Writer.Status(),
			}
			file, _ := os.OpenFile("./log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			defer file.Close()

			events, _ := json.Marshal(event)
			marshal, _ := json.Marshal(fmt.Sprintf("%v:%v", time.Now().Format(time.DateTime), string(events)))
			log.Println(string(marshal))
			file.WriteString(string(marshal) + "\n")
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type Event struct {
	UserId   interface{} `json:"userId"`
	Method   interface{} `json:"method"`
	Path     interface{} `json:"path"`
	Duration interface{} `json:"duration"`
	Status   interface{} `json:"status"`
}
