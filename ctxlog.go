package ctxlog

import (
	"context"
	"errors"
	"io"
	"log"
)

type LogContext struct {
	logger Logger
	//format LogFormatter
}

type Logger interface {
	Output(calldepth int, s string) error
}

type PriorityLogger interface {
	Logger
	Alert(string) error
	Crit(string) error
	Err(string) error
	Warning(string) error
	Notice(string) error
	Info(string) error
	Debug(string) error
}

var ErrNoLogContext = errors.New("no log object in context")

// From lambdacontext, an unexported type to be used as key for types in
// this package. This prevents collisions with keys defined in other packages.
type key struct{}

// the key for CLIContext in Contexts.
var contextKey = &key{}

func WithWriter(parent context.Context, out io.Writer) context.Context {
	return WithLogger(parent, &WrappedLogger{log.New(out, "", log.LstdFlags)})
}

func WithLogger(parent context.Context, l Logger) context.Context {
	return WithLogContext(parent, &LogContext{l})
}

func WithLogContext(parent context.Context, logctx *LogContext) context.Context {
	return context.WithValue(parent, contextKey, logctx)
}

func FromContext(ctx context.Context) (*LogContext, bool) {
	logctx, ok := ctx.Value(contextKey).(*LogContext)
	return logctx, ok
}

func Log(ctx context.Context) Logger {
	logctx, ok := FromContext(ctx)
	if !ok {
		return nil
	}
	return logctx.logger
}

func PLog(ctx context.Context) PriorityLogger {
	return Log(ctx).(PriorityLogger)
}
