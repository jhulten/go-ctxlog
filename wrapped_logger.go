package ctxlog

import "fmt"

type WrappedLogger struct {
	dest Logger
}

func (w *WrappedLogger) Emerg(m string) error {
	err := w.formatOutput(LOG_EMERG, m)
	return err
}

// Alert logs a message with severity LOG_ALERT, ignoring the severity
// passed to New.
func (w *WrappedLogger) Alert(m string) error {
	err := w.formatOutput(LOG_ALERT, m)
	return err
}

// Crit logs a message with severity LOG_CRIT, ignoring the severity
// passed to New.
func (w *WrappedLogger) Crit(m string) error {
	err := w.formatOutput(LOG_CRIT, m)
	return err
}

// Err logs a message with severity LOG_ERR, ignoring the severity
// passed to New.
func (w *WrappedLogger) Err(m string) error {
	err := w.formatOutput(LOG_ERR, m)
	return err
}

// Warning logs a message with severity LOG_WARNING, ignoring the
// severity passed to New.
func (w *WrappedLogger) Warning(m string) error {
	err := w.formatOutput(LOG_WARNING, m)
	return err
}

// Notice logs a message with severity LOG_NOTICE, ignoring the
// severity passed to New.
func (w *WrappedLogger) Notice(m string) error {
	err := w.formatOutput(LOG_NOTICE, m)
	return err
}

// Info logs a message with severity LOG_INFO, ignoring the severity
// passed to New.
func (w *WrappedLogger) Info(m string) error {
	err := w.formatOutput(LOG_INFO, m)
	return err
}

// Debug logs a message with severity LOG_DEBUG, ignoring the severity
// passed to New.
func (w *WrappedLogger) Debug(m string) error {
	err := w.formatOutput(LOG_DEBUG, m)
	return err
}

func (w *WrappedLogger) formatOutput(p Priority, m string) error {
	return w.Output(3, fmt.Sprintf("%s: %s", p, m))
}

func (w *WrappedLogger) Output(calldepth int, m string) error {
	return w.dest.Output(calldepth+1, m)
}
