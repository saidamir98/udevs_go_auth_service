package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field ...
type Field = zapcore.Field

var (
	// Int ..
	Int = zap.Int
	// String ...
	String = zap.String
	// Error ...
	Error = zap.Error
	// Bool ...
	Bool = zap.Bool
	// Any ...
	Any = zap.Any
)

// Logger ...
type LoggerI interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	DPanic(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type loggerImpl struct {
	zap *zap.Logger
}

// NewLogger ...
func NewLogger(namespace string, level string) LoggerI {
	if level == "" {
		level = LevelInfo
	}

	logger := loggerImpl{
		zap: newZapLogger(namespace, level),
	}

	return &logger
}

func (l *loggerImpl) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, fields...)
}

func (l *loggerImpl) Info(msg string, fields ...Field) {
	l.zap.Info(msg, fields...)
}

func (l *loggerImpl) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, fields...)
}

func (l *loggerImpl) Error(msg string, fields ...Field) {
	l.zap.Error(msg, fields...)
}

func (l *loggerImpl) DPanic(msg string, fields ...Field) {
	l.zap.DPanic(msg, fields...)
}

func (l *loggerImpl) Panic(msg string, fields ...Field) {
	l.zap.Panic(msg, fields...)
}

func (l *loggerImpl) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, fields...)
}

// GetNamed ...
func GetNamed(l LoggerI, name string) LoggerI {
	switch v := l.(type) {
	case *loggerImpl:
		v.zap = v.zap.Named(name)
		return v
	default:
		l.Info("logger.GetNamed: invalid logger type")
		return l
	}
}

// WithFields ...
func WithFields(l LoggerI, fields ...Field) LoggerI {
	switch v := l.(type) {
	case *loggerImpl:
		return &loggerImpl{
			zap: v.zap.With(fields...),
		}
	default:
		l.Info("logger.WithFields: invalid logger type")
		return l
	}
}

// Cleanup ...
func Cleanup(l LoggerI) error {
	switch v := l.(type) {
	case *loggerImpl:
		return v.zap.Sync()
	default:
		l.Info("logger.Cleanup: invalid logger type")
		return nil
	}
}
