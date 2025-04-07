package logger

import (
	"fmt"
	"time"

	sentrysdk "github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sentry struct {
	zapcore.Core
	fields []zapcore.Field
}

func withSentry(dsn string) zap.Option {
	return zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		err := sentrysdk.Init(sentrysdk.ClientOptions{
			Dsn:              dsn,
			EnableTracing:    true,
			TracesSampleRate: 1.0,
		})

		if err != nil {
			panic(fmt.Errorf("error initializing Sentry: %w", err))
		}

		s := sentry{Core: core, fields: nil}

		return zapcore.NewTee(core, s)
	})
}

func (s sentry) Write(entry zapcore.Entry, fields []zapcore.Field) error { //nolint:gocritic // Zap interface compliance
	m := make(map[string]any, len(fields))

	s.fields = append(s.fields, fields...)

	enc := zapcore.NewMapObjectEncoder()

	for _, f := range s.fields {
		f.AddTo(enc)
	}

	for k, v := range enc.Fields {
		m[k] = v
	}

	hub := sentrysdk.CurrentHub().Clone()

	if shouldCaptureEvent(entry.Level) { // Event
		event := sentrysdk.NewEvent()

		event.Level = zapToSentryLevel(entry.Level)
		event.Message = entry.Message
		event.Timestamp = entry.Time
		event.Tags["logger"] = entry.LoggerName
		event.Tags["stack"] = entry.Stack
		event.Tags["caller"] = entry.Caller.String()
		event.Extra = m

		hub.CaptureEvent(event)
	} else { // Breadcrumb
		hint := sentrysdk.BreadcrumbHint{
			"logger": entry.LoggerName,
			"stack":  entry.Stack,
			"caller": entry.Caller.String(),
		}

		hub.AddBreadcrumb(&sentrysdk.Breadcrumb{
			Level:     zapToSentryLevel(entry.Level),
			Timestamp: entry.Time,
			Message:   entry.Message,
			Data:      m,
		}, &hint)
	}

	return nil
}

func (s sentry) With(fields []zapcore.Field) zapcore.Core {
	return sentry{s.Core.With(fields), fields}
}

func (s sentry) Check(entry zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry { //nolint:gocritic // Zap interface compliance
	if s.Enabled(entry.Level) {
		return ce.AddCore(entry, s)
	}

	return ce
}

func (s sentry) Sync() error {
	sentrysdk.Flush(2 * time.Second)

	return nil
}

func zapToSentryLevel(l zapcore.Level) sentrysdk.Level {
	switch l { //nolint:exhaustive,nolintlint // This is OK
	case zapcore.DebugLevel:
		return sentrysdk.LevelDebug
	case zapcore.InfoLevel:
		return sentrysdk.LevelInfo
	case zapcore.WarnLevel:
		return sentrysdk.LevelWarning
	case zapcore.ErrorLevel:
		return sentrysdk.LevelError
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		return sentrysdk.LevelFatal
	default:
		return sentrysdk.LevelInfo
	}
}

func shouldCaptureEvent(level zapcore.Level) bool {
	return level >= zapcore.WarnLevel
}
