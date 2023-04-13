package magic

import (
	"darvaza.org/slog"
)

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
