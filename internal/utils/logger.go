package utils

import "github.com/phuslu/log"

var Logger = NewLogger()

func NewLogger() *log.Logger {
	logger := &log.Logger{
		Level:      log.TraceLevel,
		Caller:     1,
		TimeFormat: "2006-01-02 15:04:05",
		Writer: &log.MultiWriter{
			InfoWriter:    &log.FileWriter{Filename: "logs/cli.log", MaxSize: 100 << 20, LocalTime: false},
			ConsoleWriter: &log.ConsoleWriter{ColorOutput: true},
			ConsoleLevel:  log.InfoLevel,
		},
	}
	return logger
}
