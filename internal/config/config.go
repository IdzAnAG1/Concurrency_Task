package config

import (
	"go.uber.org/zap"
	"time"
)

type Config struct {
	FileVerifier struct {
		PathToMethodsDirectory string
		Interval               time.Duration
	}
	Logger struct {
		zap.Config
	}
}
