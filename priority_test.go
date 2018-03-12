package ctxlog

import "testing"

func TestPriority_String(t *testing.T) {
	tests := []struct {
		name string
		p    Priority
		want string
	}{
		{"LOG_EMERG", LOG_EMERG, "emergency"},
		{"LOG_ALERT", LOG_ALERT, "alert"},
		{"LOG_CRIT", LOG_CRIT, "critical"},
		{"LOG_ERR", LOG_ERR, "error"},
		{"LOG_WARNING", LOG_WARNING, "warning"},
		{"LOG_NOTICE", LOG_NOTICE, "notice"},
		{"LOG_INFO", LOG_INFO, "info"},
		{"LOG_DEBUG", LOG_DEBUG, "debug"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Priority.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
