package logging

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

// GenerateLogInfo generates base log information
func GenerateLogInfo(c *gin.Context, start time.Time) LogInfo {
	return LogInfo{
		ClientIP:    c.ClientIP(),
		Date:        start.Format(time.RFC3339),
		Method:      c.Request.Method,
		RequestURI:  c.Request.URL.RequestURI(),
		Referer:     c.Request.Referer(),
		HTTPVersion: c.Request.Proto,
		Size:        c.Writer.Size(),
		Status:      c.Writer.Status(),
		UserAgent:   c.Request.UserAgent(),
		Latency:     time.Now().Sub(start),
	}
}

// ConvertToMapFromBody converts to a map from a request body
func ConvertToMapFromBody(c *gin.Context) (string, error) {
	s := ""
	b, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println("byte array converted to string:", string(b))
	fmt.Println("error:", err)
	if err != nil {
		return s, err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	if len(b) != 0 {
		s = string(b)
	}
	return s, nil
}
