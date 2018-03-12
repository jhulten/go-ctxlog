package ctxlog

type Priority int

const (
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

var priorityString = []string{"emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"}

func (p Priority) String() string {
	return priorityString[p]
}
