package config

import (
	"io"
	"os"

	"github.com/imdario/mergo"
	"github.com/rs/zerolog"
)

type LogConfig struct {
	// io.Writer to which to write the logs (defaults to os.stdout)
	Out io.Writer
	// log level for loggers
	Level zerolog.Level
}

func createDefaultLogConfig(config LogConfig) LogConfig {
	cfg := LogConfig{
		Out:   os.Stderr,
		Level: zerolog.DebugLevel,
	}

	mergo.Merge(&config, cfg)

	return config
}
