package console

import (
	"fmt"
	"time"
)

const (
	// LogLevelDebug DEBUG
	LogLevelDebug = 0

	// LogLevelInfo INFO
	LogLevelInfo = 1

	// LogLevelWarning WARNING
	LogLevelWarning = 2

	// LogLevelError ERROR
	LogLevelError = 3
)

// LogEntry an entry of log
type LogEntry struct {
	Context interface{}
	Logger  string
	At      time.Time
	Level   int
	Message string
}

var logChannel chan LogEntry
var isDebugEnabled bool

func init() {

	// debug enabled if environment is dev
	// isDebugEnabled = config.Cfg.GetString("environment") == "dev"
	isDebugEnabled = false

	// TODO: use logger.conf to define the logger level.

	logChannel = make(chan LogEntry, 256)

	formats := []string{
		"\033[0;36m%s DBG [%s] %s\033[0m\n", // debug
		"%s INF [%s] %s\n",                  // info
		"\033[1;33m%s WRN [%s] %s\033[0m\n", // warning
		"\033[1;31m%s ERR [%s] %s\033[0m\n", // error
	}

	go func() {
		for {
			entry := <-logChannel

			// TODO: check for logger level
			fmt.Printf(formats[entry.Level], entry.At.Format("2006-01-02T15:04:05.000Z07:00"), entry.Logger, entry.Message)
		}
	}()
}

// Logger logger
type Logger struct {
	Name string
}

// GetLogger by given name
func GetLogger(name string) *Logger {
	return &Logger{
		Name: name,
	}
}

// IsDebugEnabled true if environment is "dev"
func (a *Logger) IsDebugEnabled() bool {
	return isDebugEnabled
}

// Debug print a debug message
func (a *Logger) Debug(message string) {
	if a.IsDebugEnabled() {
		Debug(a.Name, message)
	}
}

// DebugF print a debug message with formatting
func (a *Logger) DebugF(message string, params ...interface{}) {
	if a.IsDebugEnabled() {
		DebugF(a.Name, message, params...)
	}
}

// Info print an info message
func (a *Logger) Info(message string) {
	Info(a.Name, message)
}

// InfoF print an info message
func (a *Logger) InfoF(message string, params ...interface{}) {
	InfoF(a.Name, message, params...)
}

// Warning print a warning message
func (a *Logger) Warning(message string) {
	Warning(a.Name, message)
}

// WarningF print a warning message
func (a *Logger) WarningF(message string, params ...interface{}) {
	WarningF(a.Name, message, params...)
}

// Error print an error message
func (a *Logger) Error(message string) {
	Error(a.Name, message)
}

// ErrorF print an error message
func (a *Logger) ErrorF(message string, params ...interface{}) {
	ErrorF(a.Name, message, params...)
}

// Debug print a debug message
func Debug(logger string, message string) {
	if isDebugEnabled {
		logChannel <- LogEntry{
			Logger:  logger,
			At:      time.Now(),
			Level:   LogLevelDebug,
			Message: message,
		}
	}
}

// DebugF print a debug message with formatting
func DebugF(logger string, message string, params ...interface{}) {
	if isDebugEnabled {
		logChannel <- LogEntry{
			Logger:  logger,
			At:      time.Now(),
			Level:   LogLevelDebug,
			Message: fmt.Sprintf(message, params...),
		}
	}
}

// Info print an info message
func Info(logger string, message string, params ...interface{}) {
	logChannel <- LogEntry{
		Logger:  logger,
		At:      time.Now(),
		Level:   LogLevelInfo,
		Message: message,
	}
}

// InfoF print an info message with formatting
func InfoF(logger string, message string, params ...interface{}) {
	logChannel <- LogEntry{
		Logger:  logger,
		At:      time.Now(),
		Level:   LogLevelInfo,
		Message: fmt.Sprintf(message, params...),
	}
}

// Warning print a warning message
func Warning(logger string, message string, params ...interface{}) {
	logChannel <- LogEntry{
		Logger:  logger,
		At:      time.Now(),
		Level:   LogLevelWarning,
		Message: message,
	}
}

// WarningF print a warning message with formatting
func WarningF(logger string, message string, params ...interface{}) {
	logChannel <- LogEntry{
		Logger:  logger,
		At:      time.Now(),
		Level:   LogLevelWarning,
		Message: fmt.Sprintf(message, params...),
	}
}

// Error print an error message
func Error(logger string, message string) {
	logChannel <- LogEntry{
		Logger:  logger,
		At:      time.Now(),
		Level:   LogLevelError,
		Message: message,
	}
}

// ErrorF print an error message with formatting
func ErrorF(logger string, message string, params ...interface{}) {
	logChannel <- LogEntry{
		Logger:  logger,
		At:      time.Now(),
		Level:   LogLevelError,
		Message: fmt.Sprintf(message, params...),
	}
}
