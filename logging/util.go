package logging

import (
	"bytes"
	"io/ioutil"
	"time"

	"encoding/json"
	"github.com/gin-gonic/gin"
)

func generateLogInfo(c *gin.Context, start time.Time) logInfo {
	return logInfo{
		ClientIP:    c.ClientIP(),
		Date:        start.Format(time.RFC3339),
		Method:      c.Request.Method,
		Path:        c.Request.URL.Path,
		RawQuery:    c.Request.URL.RawQuery,
		HTTPVersion: c.Request.Proto,
		Size:        c.Writer.Size(),
		Status:      c.Writer.Status(),
		UserAgent:   c.Request.UserAgent(),
		Latency:     time.Now().Sub(start),
	}
}

func convertToMapFromBody(c *gin.Context) (m map[string]interface{}, err error) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	err = json.Unmarshal(b, &m)
	return
}
