package grpcgateway

import (
	"github.com/phuslu/log"
)

// New wraps the Logger to provide a GrpcGateway logger
func New(l *log.Logger) logger {
	l1 := *l
	switch {
	case l1.Caller > 0:
		l1.Caller++
	case l1.Caller < 0:
		l1.Caller--
	}
	return logger{logger: l1}
}

// logger implements methods to satisfy interface
// github.com/grpc-ecosystem/go-grpc-middleware/blob/v2/interceptors/logging/logging.go
type logger struct {
	logger log.Logger
}

// WithValues adds some key-value pairs of context to a logger.
// See Info for documentation on how key/value pairs work.
func (g logger) WithValues(keysAndValues ...interface{}) logger {
	l := g.logger
	l.Context = append(g.logger.Context, log.NewContext(nil).KeysAndValues(keysAndValues...).Value()...)
	return logger{logger: l}
}

// Debug logs a debug with the message and key/value pairs as context.
func (g logger) Debug(msg string) {
	g.logger.Debug().Msg(msg)
}

// Info logs an info with the message and key/value pairs as context.
func (g logger) Info(msg string) {
	g.logger.Info().Msg(msg)
}

// Warning logs a warning with the message and key/value pairs as context.
func (g logger) Warning(msg string) {
	g.logger.Warn().Msg(msg)
}

// Error logs an error with the message and key/value pairs as context.
func (g logger) Error(msg string) {
	g.logger.Error().Msg(msg)
}
