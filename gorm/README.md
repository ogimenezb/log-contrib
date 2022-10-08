# gorm

[![Go Reference](https://pkg.go.dev/badge/github.com/phuslu/log-contrib/gorm.svg)](https://pkg.go.dev/github.com/phuslu/log-contrib/gorm)
[![Go Report Card](https://goreportcard.com/badge/github.com/phuslu/log-contrib/gorm)](https://goreportcard.com/report/github.com/phuslu/log-contrib/gorm)

A [gorm](https://github.com/go-gorm/gorm) logger implementation.

## Usage

```go
import (
    "github.com/mattn/go-sqlite3"
    "github.com/phuslu/log"
    "gorm.io/gorm"
    gormlogger "github.com/phuslu/log-contrib/gorm"
)

func main() {
    newLogger := gormlogger.New(&log.Logger{
        Level:      log.InfoLevel,
        TimeFormat: "15:04:05",
        Caller:     1,
        Writer: &log.ConsoleWriter{
            ColorOutput:    true,
            QuoteString:    false,
            EndWithMessage: false,
        },
    })

    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: newLogger})

    // ...
}
```
