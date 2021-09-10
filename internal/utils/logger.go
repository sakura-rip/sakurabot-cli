package utils

import "github.com/phuslu/log"

var DefaultLogger = NewLogger()

func NewLogger() *log.Logger {
	logger := &log.Logger{
		Level:      log.TraceLevel,
		TimeFormat: "01-02 15:04:05",
		Writer: &log.MultiWriter{
			InfoWriter:    &log.FileWriter{Filename: "logs/cli.log", MaxSize: 100 << 20, LocalTime: false},
			ConsoleWriter: &log.ConsoleWriter{ColorOutput: true},
			ConsoleLevel:  log.InfoLevel,
		},
	}
	return logger
}

// Info starts a new message with info level.
func Info() *log.Entry {
	return DefaultLogger.Info()
}

// Warn starts a new message with warning level.
func Warn() *log.Entry {
	return DefaultLogger.Warn()
}

// Error starts a new message with error level.
func Error() *log.Entry {
	return DefaultLogger.Error()
}

// Fatal starts a new message with fatal level.
func Fatal() *log.Entry {
	return DefaultLogger.Fatal()
}

// Panic starts a new message with panic level.
func Panic() *log.Entry {
	return DefaultLogger.Panic()
}

// Log starts a new message with no level.
func Log() *log.Entry {
	return DefaultLogger.Log()
}
