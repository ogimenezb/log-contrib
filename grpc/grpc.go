package grpc

import (
	"github.com/phuslu/log"
)

// New wraps the log.Logger to provide a LoggerV2 logger
func New(l *log.Logger) (g *logger) {
	l1 := *l
	switch {
	case l1.Caller > 0:
		l1.Caller++
	case l1.Caller < 0:
		l1.Caller--
	}
	return &logger{logger: &l1}
}

// logger implements methods to satisfy interface
// google.golang.org/grpc/grpclog.LoggerV2.
type logger struct {
	logger *log.Logger
}

// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
func (g *logger) Info(args ...interface{}) {
	g.logger.Info().Msgs(args...)
}

// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
func (g *logger) Infoln(args ...interface{}) {
	g.logger.Info().Msgs(args...)
}

// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
func (g *logger) Infof(format string, args ...interface{}) {
	g.logger.Info().Msgf(format, args...)
}

// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (g *logger) Warning(args ...interface{}) {
	g.logger.Warn().Msgs(args...)
}

// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (g *logger) Warningln(args ...interface{}) {
	g.logger.Warn().Msgs(args...)
}

// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (g *logger) Warningf(format string, args ...interface{}) {
	g.logger.Warn().Msgf(format, args...)
}

// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
func (g *logger) Error(args ...interface{}) {
	g.logger.Error().Msgs(args...)
}

// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
func (g *logger) Errorln(args ...interface{}) {
	g.logger.Error().Msgs(args...)
}

// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
func (g *logger) Errorf(format string, args ...interface{}) {
	g.logger.Error().Msgf(format, args...)
}

// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (g *logger) Fatal(args ...interface{}) {
	g.logger.Fatal().Msgs(args...)
}

// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (g *logger) Fatalln(args ...interface{}) {
	g.logger.Fatal().Msgs(args...)
}

// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (g *logger) Fatalf(format string, args ...interface{}) {
	g.logger.Fatal().Msgf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose leveg.
func (g *logger) V(level int) bool {
	return level >= int(g.logger.Level)
}

type grpcLoggerV2 interface {
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})
	Warning(args ...interface{})
	Warningln(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})
	V(l int) bool
}

var _ grpcLoggerV2 = (*logger)(nil)
