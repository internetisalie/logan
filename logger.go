package logan

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Logger struct {
	ParentLogger
	Fields logrus.Fields
}

func (logger *Logger) WithField(key string, value interface{}) *logrus.Entry {
	return logger.ParentLogger.WithFields(logger.Fields).WithField(key, value)
}

func (logger *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.ParentLogger.WithFields(logger.Fields).WithFields(fields)
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func (logger *Logger) WithError(err error) *logrus.Entry {
	return logger.ParentLogger.WithFields(logger.Fields).WithError(err)
}

// Add a context to the log entry.
func (logger *Logger) WithContext(ctx context.Context) *logrus.Entry {
	return logger.ParentLogger.WithFields(logger.Fields).WithContext(ctx)
}

// Overrides the time of the log entry.
func (logger *Logger) WithTime(t time.Time) *logrus.Entry {
	return logger.ParentLogger.WithFields(logger.Fields).WithTime(t)
}

func (logger *Logger) WithExtendedField(key string, value interface{}) *Logger {
	return newLogger(logger, logrus.Fields{key:value})
}

func (logger *Logger) WithExtendedFields(fields ...logrus.Fields) *Logger {
	if len(fields) == 0 {
		return logger
	}

	return newLogger(logger, fields...)
}

func (logger *Logger) Logf(level logrus.Level, format string, args ...interface{}) {
	if logger.ParentLogger.IsLevelEnabled(level) {
		logger.ParentLogger.WithFields(logger.Fields).Logf(level, format, args...)
	}
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Tracef(format, args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Debugf(format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Infof(format, args...)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Printf(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Warnf(format, args...)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Warningf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Errorf(format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Fatalf(format, args...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Panicf(format, args...)
}

func (logger *Logger) Log(level logrus.Level, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Log(level, args...)
}

func (logger *Logger) Trace(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Trace(args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Debug(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Info(args...)
}

func (logger *Logger) Print(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Print(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Warn(args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Warning(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Error(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Fatal(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Panic(args...)
}

func (logger *Logger) Logln(level logrus.Level, args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Logln(level, args...)
}

func (logger *Logger) Traceln(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Traceln(args...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Debugln(args...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Infoln(args...)
}

func (logger *Logger) Println(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Println(args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Warnln(args...)
}

func (logger *Logger) Warningln(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Warningln(args...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Errorln(args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Fatalln(args...)
}

func (logger *Logger) Panicln(args ...interface{}) {
	logger.ParentLogger.WithFields(logger.Fields).Panicln(args...)
}

func newLogger(logger ParentLogger, fields ...logrus.Fields) *Logger {
	allFields := make(logrus.Fields)
	for _, field := range fields {
		for k, v := range field {
			allFields[k] = v
		}
	}

	return &Logger{
		ParentLogger: logger,
		Fields:       allFields,
	}
}

func NewLogger(name string, fields ...logrus.Fields) *Logger {
	fields = append([]logrus.Fields{{"name": name}}, fields...)
	return newLogger(logrus.StandardLogger(), fields...)
}
