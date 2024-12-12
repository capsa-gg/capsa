package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogConfigDev(encodeLevel zapcore.LevelEncoder, output string) zap.Config {
	return zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableStacktrace: true, // For now we don't want full stacktraces for everything
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "message",
			LevelKey:         "level",
			TimeKey:          "time",
			NameKey:          "name",
			CallerKey:        "caller",
			StacktraceKey:    "stacktrace",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      encodeLevel, // To not litter the .log output with console coloring, change to zapcore.CapitalLevelEncoder
			EncodeTime:       zapcore.TimeEncoderOfLayout("15:04:05.000"),
			EncodeDuration:   zapcore.StringDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			ConsoleSeparator: "\t",
		},
		OutputPaths:      []string{output},
		ErrorOutputPaths: []string{output},
	}
}

func newLogConfigProd(_ zapcore.LevelEncoder, _ string) zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "message",
			LevelKey:         "level",
			TimeKey:          "time",
			NameKey:          "name",
			CallerKey:        "caller",
			FunctionKey:      "function",
			StacktraceKey:    "stacktrace",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      zapcore.CapitalLevelEncoder,
			EncodeTime:       zapcore.ISO8601TimeEncoder,
			EncodeDuration:   zapcore.StringDurationEncoder,
			EncodeCaller:     zapcore.FullCallerEncoder,
			EncodeName:       zapcore.FullNameEncoder,
			ConsoleSeparator: "\t",
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    nil,
	}
}

func generateConfig(isDev bool) zap.Config {
	// DEVELOPMENT
	if isDev {
		return newLogConfigDev(zapcore.CapitalColorLevelEncoder, "stderr")
	}

	// PRODUCTION
	return newLogConfigProd(zapcore.CapitalColorLevelEncoder, "stderr")
}

// New returns a *zap.Logger for the given environment.
// If the sentryDns argument is not an empty string, Sentry will be initialized and attached to the logger.
func New(isDev bool, sentryDsn string) (*zap.Logger, error) {
	config := generateConfig(isDev)
	loggerBase, err := config.Build()

	if err != nil {
		return nil, fmt.Errorf("error building logging config: %w", err)
	}

	core := zapcore.NewTee(loggerBase.Core())

	if sentryDsn != "" {
		logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), withSentry(sentryDsn))

		return logger, nil
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger, nil
}
