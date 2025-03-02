package logger

import "go.uber.org/zap"

var zl *zap.Logger

func init() {
	zapLogger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(zapLogger)
	defer zl.Sync()
	zl = zapLogger
}

type LogField[T any] struct {
	Key   string
	Value T
}

func (l LogField[T]) toZapField() zap.Field {
	switch any(l.Value).(type) {
	case int:
		return zap.Int(l.Key, any(l.Value).(int))
	case string:
		return zap.String(l.Key, any(l.Value).(string))
	case bool:
		return zap.Bool(l.Key, any(l.Value).(bool))
	default:
		return zap.Any(l.Key, l.Value)
	}
}

func Debug[T any](message string, logFields []LogField[T]) {
	var fields []zap.Field
	for _, field := range logFields {
		fields = append(fields, field.toZapField())
	}
	zl.Debug(message, fields...)
}

func Info[T any](message string, logFields []LogField[T]) {
	var fields []zap.Field
	for _, field := range logFields {
		fields = append(fields, field.toZapField())
	}
	zl.Info(message, fields...)
}

func Warn[T any](message string, logFields []LogField[T]) {
	var fields []zap.Field
	for _, field := range logFields {
		fields = append(fields, field.toZapField())
	}
	zl.Warn(message, fields...)
}

func Error[T any](message string, logFields []LogField[T]) {
	var fields []zap.Field
	for _, field := range logFields {
		fields = append(fields, field.toZapField())
	}
	zl.Error(message, fields...)
}
