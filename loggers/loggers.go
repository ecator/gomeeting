package loggers

import (
	"log"
	"os"
)

// Logger looger
type Logger struct {
	logger log.Logger
	file   *os.File
}

func (l *Logger) output(msg string) {
	if msg[0:5] == "ERROR" {
		l.logger.SetOutput(os.Stderr)
	} else {
		l.logger.SetOutput(os.Stdout)
	}
	l.logger.Println(msg)
	l.logger.SetOutput(l.file)
	l.logger.Println(msg)
}

// Info outputs info
func (l *Logger) Info(msg string) {
	msg = "INFO  " + msg
	l.output(msg)
}

// Error outputs error
func (l *Logger) Error(msg string) {
	msg = "ERROR " + msg
	l.output(msg)
}

// Warn outputs warn
func (l *Logger) Warn(msg string) {
	msg = "WRAN  " + msg
	l.output(msg)
}

// Close closes log file
func (l *Logger) Close() {
	l.file.Close()
}

// New returns a new logger
func New(logFile string, prefix string) (*Logger, error) {
	logger := new(Logger)
	var err error
	logger.file, err = os.OpenFile(logFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	logger.logger.SetOutput(logger.file)
	logger.logger.SetPrefix(prefix)
	logger.logger.SetFlags(log.LstdFlags)
	return logger, nil
}
