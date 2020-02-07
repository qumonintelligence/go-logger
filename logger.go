package logger

import "github.com/qumonintelligence/go-logger/v2/console"

// ILogger logger interface
type ILogger interface {
	Debug(message string)
	Info(message string)
	Warning(message string)
	Error(message string)

	DebugF(message string, params ...interface{})
	InfoF(message string, params ...interface{})
	WarningF(message string, params ...interface{})
	ErrorF(message string, params ...interface{})
}

const defaultLogger = "default"

// Loggable make a struct loggable
type Loggable struct {
	Logger ILogger
}

// DebugF print a debug message.  If Logger is not defined, a general logger is used
func (s *Loggable) DebugF(message string, params ...interface{}) {
	if s.Logger == nil {
		console.DebugF(defaultLogger, message, params...)
		return
	}

	s.Logger.DebugF(message, params...)
}

// InfoF print an info message.  If Logger is not defined, a general logger is used
func (s *Loggable) InfoF(message string, params ...interface{}) {
	if s.Logger == nil {
		console.InfoF(defaultLogger, message, params...)
		return
	}

	s.Logger.InfoF(message, params...)
}

// WarningF print an info message.  If Logger is not defined, a general logger is used
func (s *Loggable) WarningF(message string, params ...interface{}) {
	if s.Logger == nil {
		console.WarningF(defaultLogger, message, params...)
		return
	}

	s.Logger.WarningF(message, params...)
}

// ErrorF print an error message.  If Logger is not defined, a general logger is used
func (s *Loggable) ErrorF(message string, params ...interface{}) {
	if s.Logger == nil {
		console.ErrorF(defaultLogger, message, params...)
		return
	}

	s.Logger.ErrorF(message, params...)
}
