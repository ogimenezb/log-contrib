package gin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

func New(logger *log.Logger, skip func(c *gin.Context) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		if skip != nil && skip(c) {
			return
		}

		end := time.Now()
		latency := end.Sub(start)

		path := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			path = path + "?" + c.Request.URL.RawQuery
		}
		msg := "Request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}
		status := c.Writer.Status()

		var e *log.Entry
		switch {
		case status >= 400 && status < 500:
			e = logger.Warn()
		case status >= 500:
			e = logger.Error()
		default:
			e = logger.Info()
		}
		e.Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Str("user_agent", c.Request.UserAgent()).
			Msg(msg)
	}
}
