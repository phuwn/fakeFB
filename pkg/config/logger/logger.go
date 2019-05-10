package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

// https://github.com/uber-go/zap
// We use zap to log,if someone wants to replace it with another logging lib, change it here
// note that we have to import "go.uber.org/zap" instead of github

// Logger is a interface for leveled logging
// We will use template logging instead of structured logging
type Logger interface {
	Logf(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

// LoggingService is a struct implement logger interface
type LoggingService struct {
	logger *zap.SugaredLogger
}

// NewLogger return a logger instance for app
func NewLogger() *LoggingService {
	return &LoggingService{
		logger: newLogger(),
	}
}

func newLogger() *zap.SugaredLogger {
	opts := zap.AddCallerSkip(1)
	logger, err := zap.NewProduction(opts)
	if err != nil {
		log.Fatal("Can't init zap log")
	}
	defer logger.Sync()
	return logger.Sugar()
}

// Logf ...
func (l *LoggingService) Logf(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	l.logger.Infof(s)
}

// Debugf ...
func (l *LoggingService) Debugf(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	l.logger.Debugf(s)
}

// Infof ...
func (l *LoggingService) Infof(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	l.logger.Infof(s)
}

// Warnf ...
func (l *LoggingService) Warnf(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	l.logger.Warnf(s)
}

// Errorf ...
func (l *LoggingService) Errorf(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	l.logger.Errorf(s)
}
