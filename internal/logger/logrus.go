package logger

import "github.com/sirupsen/logrus"

type LogrusLogger struct{}

func NewLogrusLogger() LogrusLogger {
	return LogrusLogger{}
}

func (l LogrusLogger) Debug(msg ...interface{}) {
	logrus.Debug(msg...)
}

func (l LogrusLogger) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (l LogrusLogger) Info(msg ...interface{}) {
	logrus.Info(msg...)
}

func (l LogrusLogger) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (l LogrusLogger) Warn(msg ...interface{}) {
	logrus.Warn(msg...)
}

func (l LogrusLogger) Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func (l LogrusLogger) Error(msg ...interface{}) {
	logrus.Error(msg...)
}

func (l LogrusLogger) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
