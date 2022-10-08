# logr

[![Go Reference](https://pkg.go.dev/badge/github.com/phuslu/log-contrib/logr.svg)](https://pkg.go.dev/github.com/phuslu/log-contrib/logr)
[![Go Report Card](https://goreportcard.com/badge/github.com/phuslu/log-contrib/logr)](https://goreportcard.com/report/github.com/phuslu/log-contrib/logr)

A [logr](https://github.com/go-logr/logr) LogSink implementation.

## Usage

```go
import (
    "github.com/phuslu/log-contrib/logr"
    "github.com/phuslu/log"
)

func main() {
    logger := logr.New(&log.Logger{
        Level:      log.InfoLevel,
        TimeFormat: "15:04:05",
        Caller:     1,
        Writer: &log.ConsoleWriter{
            ColorOutput:    true,
            QuoteString:    false,
            EndWithMessage: false,
        },
    })

    logger.Info("Logr in action!", "the answer", 42)
}
```

## Implementation Details

For the most part, concepts in Phuslog correspond directly with those in log/logr.

Levels in log/logr correspond to custom debug levels in Phuslog. Any given level
in log/logr is represents by `phuslogLevel = 1 - logrLevel`.

For example `V(2)` is equivalent to Phuslog's `TraceLevel`, while `V(1)` is
equivalent to Phuslog's `DebugLevel`.
