package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(namespace, level, timeFormat string) *zap.Logger {
	globalLevel := parseLevel(level)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= globalLevel && lvl < zapcore.ErrorLevel
	})

	consoleInfos := zapcore.Lock(os.Stdout)

	consoleErrors := zapcore.Lock(os.Stderr)

	// Configure console output.
	encoderCfg := zap.NewProductionEncoderConfig()
	if len(timeFormat) > 0 {
		encoderCfg.EncodeTime = customTimeEncoder(timeFormat)
	} else {
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	}
	consoleEncoder := zapcore.NewJSONEncoder(encoderCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleInfos, lowPriority),
	)

	logger := zap.New(core)

	logger = logger.Named(namespace)

	zap.RedirectStdLog(logger)

	return logger
}

func customTimeEncoder(timeFormat string) func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(timeFormat))
	}
}

func parseLevel(level string) zapcore.Level {
	switch level {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	case LevelDPanic:
		return zapcore.DPanicLevel
	case LevelPanic:
		return zapcore.PanicLevel
	case LevelFatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
