package logr

import (
	"github.com/go-logr/logr"
	"github.com/phuslu/log"
)

// New returns a new logr.Logger instance.
func New(logger *log.Logger) logr.Logger {
	l := *logger
	return logr.New(logSink{"", nil, &l})
}

type logSink struct {
	Name    string
	Context log.Context
	Logger  *log.Logger
}

// Init receives optional information about the logr library for logSink
// implementations that need it.
func (ls logSink) Init(info logr.RuntimeInfo) {
	switch {
	case ls.Logger.Caller > 0:
		ls.Logger.Caller += 3
	case ls.Logger.Caller < 0:
		ls.Logger.Caller -= 3
	}
}

// Enabled tests whether this logSink is enabled at the specified V-level.
// For example, commandline flags might be used to set the logging
// verbosity and disable some info logs.
func (ls logSink) Enabled(level int) bool {
	const traceLevel = 1 - int(log.TraceLevel)
	return level <= traceLevel
}

// Info logs a non-error message with the given key/value pairs as context.
// The level argument is provided for optional logging.  This method will
// only be called when Enabled(level) is true. See Logger.Info for more
// detaisink.
func (ls logSink) Info(level int, msg string, keysAndValues ...interface{}) {
	if !ls.Enabled(level) {
		return
	}

	e := ls.Logger.Info()
	if ls.Name != "" {
		e.Str("logger", ls.Name)
	}
	e.Context(ls.Context).KeysAndValues(keysAndValues...).Msg(msg)
}

// Error logs an error, with the given message and key/value pairs as
// context.  See Logger.Error for more detaisink.
func (ls logSink) Error(err error, msg string, keysAndValues ...interface{}) {
	e := ls.Logger.Error()
	if ls.Name != "" {
		e.Str("logger", ls.Name)
	}
	e.Context(ls.Context).Err(err).KeysAndValues(keysAndValues...).Msg(msg)
}

// WithValues returns a new logSink with additional key/value pairs.  See
// Logger.WithValues for more detaisink.
func (ls logSink) WithValues(keysAndValues ...interface{}) logr.LogSink {
	ls.Context = log.NewContext(ls.Context[:]).KeysAndValues(keysAndValues...).Value()
	return ls
}

// WithName returns a new logSink with the specified name appended.  See
// Logger.WithName for more detaisink.
func (ls logSink) WithName(name string) logr.LogSink {
	if ls.Name == "" {
		ls.Name = name
	} else {
		ls.Name += "/" + name
	}
	return ls
}
