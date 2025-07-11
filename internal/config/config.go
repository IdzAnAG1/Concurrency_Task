package config

import (
	"time"
)

type Config struct {
	PathToMethodsDirectory string
	Interval               time.Duration
}
