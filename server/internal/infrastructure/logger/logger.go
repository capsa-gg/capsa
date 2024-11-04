package logger

import (
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
func New(isDev bool) (*zap.Logger, error) {
	config := generateConfig(isDev)
	loggerBase, err := config.Build()

	core := zapcore.NewTee(loggerBase.Core())
	logger := zap.New(core, zap.AddCaller())

	return logger, err
}
