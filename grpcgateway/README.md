# grpcgateway

[![Go Reference](https://pkg.go.dev/badge/github.com/phuslu/log-contrib/grpcgateway.svg)](https://pkg.go.dev/github.com/phuslu/log-contrib/grpcgateway)
[![Go Report Card](https://goreportcard.com/badge/github.com/phuslu/log-contrib/grpcgateway)](https://goreportcard.com/report/github.com/phuslu/log-contrib/grpcgateway)

A [grpcgateway](google.golang.org/grpcgateway/grpcgatewaylog.LoggerV2) LoggerV2 implementation.

## Usage

```go
import (
    "github.com/phuslu/log-contrib/grpcgateway"
    "github.com/phuslu/log"
)

func main() {
    logger := grpcgateway.New(&log.Logger{
        Level:      log.InfoLevel,
        TimeFormat: "15:04:05",
        Caller:     1,
        Writer: &log.ConsoleWriter{
            ColorOutput:    true,
            QuoteString:    false,
            EndWithMessage: false,
        },
    })

    logger.WithValues("foo", "bar").Info("GrpcGateway in action!")
}
```
