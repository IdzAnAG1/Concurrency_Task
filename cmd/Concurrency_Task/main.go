package main

import (
	"concurrency_task/internal/config"
	"concurrency_task/internal/file_verifier"
	_ "concurrency_task/internal/tasks/task_impl"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	zCfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:          "MSG",
			LevelKey:            "LEVEL",
			TimeKey:             "TIME",
			NameKey:             "LOG",
			CallerKey:           "CALL",
			FunctionKey:         "FUNC",
			StacktraceKey:       "ST",
			SkipLineEnding:      false,
			LineEnding:          zapcore.DefaultLineEnding,
			EncodeLevel:         zapcore.CapitalColorLevelEncoder,
			EncodeTime:          zapcore.ISO8601TimeEncoder,
			EncodeDuration:      zapcore.StringDurationEncoder,
			EncodeCaller:        zapcore.FullCallerEncoder,
			EncodeName:          zapcore.FullNameEncoder,
			NewReflectedEncoder: nil,
			ConsoleSeparator:    "~",
		},
		OutputPaths:      []string{"stdout", "dev_app.log"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    nil,
	}
	Cfg := config.Config{
		FileVerifier: struct {
			PathToMethodsDirectory string
			Interval               time.Duration
		}{
			PathToMethodsDirectory: "internal/tasks/task_impl",
			Interval:               250 * time.Millisecond,
		},
		Logger: struct {
			zap.Config
		}{zCfg},
	}
	v := file_verifier.NewVerifier(Cfg)
	v.Run()
}
