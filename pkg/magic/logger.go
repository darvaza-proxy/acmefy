package magic

import (
	"darvaza.org/slog"
	"darvaza.org/slog/handlers/discard"
)

func defaultLogger() slog.Logger {
	return discard.New()
}

// SetLogger attaches an [slog.Logger] to the client.
// if nil a [discard.Logger] will be used.
func (m *Magic) SetLogger(l slog.Logger) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if l == nil {
		l = defaultLogger()
	}
	m.logger = l
}
