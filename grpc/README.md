# grpc

[![Go Reference](https://pkg.go.dev/badge/github.com/phuslu/log-contrib/grpc.svg)](https://pkg.go.dev/github.com/phuslu/log-contrib/grpc)
[![Go Report Card](https://goreportcard.com/badge/github.com/phuslu/log-contrib/grpc)](https://goreportcard.com/report/github.com/phuslu/log-contrib/grpc)

A [grpc](google.golang.org/grpc/grpclog.LoggerV2) LoggerV2 implementation.

## Usage

```go
import (
    "github.com/phuslu/log-contrib/grpc"
    "github.com/phuslu/log"
)

func main() {
    logger := grpc.New(&log.Logger{
        Level:      log.InfoLevel,
        TimeFormat: "15:04:05",
        Caller:     1,
        Writer: &log.ConsoleWriter{
            ColorOutput:    true,
            QuoteString:    false,
            EndWithMessage: false,
        },
    })

    logger.Info("Grpc in action!", "the answer", 42)
}
```
