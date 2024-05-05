## gin

[![Go Reference](https://pkg.go.dev/badge/github.com/phuslu/log-contrib/gin.svg)](https://pkg.go.dev/github.com/phuslu/log-contrib/gin)
[![Go Report Card](https://goreportcard.com/badge/github.com/phuslu/log-contrib/gin)](https://goreportcard.com/report/github.com/phuslu/log-contrib/gin)

A [gin](https://github.com/gin-gonic/gin) logger middleware implementation.

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
	ginlogger "github.com/phuslu/log-contrib/gin"
)

func main() {
	if log.IsTerminal(os.Stderr.Fd()) {
		log.DefaultLogger = log.Logger{
			TimeFormat: "15:04:05",
			Caller:     1,
			Writer: &log.ConsoleWriter{
				ColorOutput:    true,
				QuoteString:    true,
				EndWithMessage: true,
			},
		}
	}

	if gin.IsDebugging() {
		log.DefaultLogger.SetLevel(log.DebugLevel)
	}

	r := gin.New()

	// Custom logger
	r.Use(ginlogger.New(
		&log.Logger{
			Context: log.NewContext(nil).Str("foo", "bar").Value(),
			Writer: &log.FileWriter{
				Filename: "access.log",
				MaxSize:  1024 * 1024 * 1024,
			},
		},
		func(c *gin.Context) bool {
			if c.Request.URL.Path == "/backdoor" {
				return true
			}
			return false
		}),
	)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	r.GET("/backdoor", func(c *gin.Context) {
		c.String(http.StatusOK, "a backdoor, go away")
	})

	r.Run(":8080")
}
```
