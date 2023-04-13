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

func (m *Magic) error(err error) slog.Logger {
	l := m.logger.Error()
	if err != nil {
		l = l.WithField(slog.ErrorFieldName, err)
	}
	return l
}

func (m *Magic) warn(err error) slog.Logger {
	l := m.logger.Warn()
	if err != nil {
		l = l.WithField(slog.ErrorFieldName, err)
	}
	return l
}

func (m *Magic) info() slog.Logger {
	return m.logger.Info()
}

func (m *Magic) debug() slog.Logger {
	return m.logger.Debug()
}

func (m *Magic) withInfo() (slog.Logger, bool) {
	return m.logger.Info().WithEnabled()
}

func (m *Magic) withDebug() (slog.Logger, bool) {
	return m.logger.Debug().WithEnabled()
}
