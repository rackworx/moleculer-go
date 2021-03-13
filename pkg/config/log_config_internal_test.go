package config

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestDefaultLogConfig(t *testing.T) {
	assert.Equal(t, os.Stderr, createDefaultLogConfig(LogConfig{}).Out)
	assert.Equal(t, zerolog.DebugLevel, createDefaultLogConfig(LogConfig{}).Level)
	assert.Equal(t, zerolog.InfoLevel, createDefaultLogConfig(LogConfig{
		Level: zerolog.InfoLevel,
	}).Level)
}
