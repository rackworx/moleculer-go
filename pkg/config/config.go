package config

import (
	"fmt"
	"os"
	"time"

	"github.com/rackworx/moleculer-go/pkg/errors"
)

type CheckFunc func(errors.Error) bool

type RetryPolicyConfig struct {
	Enabled  bool          //	Enable feature.
	Retries  int           //Count of retries.
	Delay    time.Duration // First delay.
	MaxDelay time.Duration // Maximum delay.
	Factor   int           // Backoff factor for delay. 2 means exponential backoff.
	Check    CheckFunc     // A function to check failed requests.
}

type TrackingConfig struct {
	Enabled         bool
	ShutdownTimeout time.Duration
}

type RegistryConfig struct {
	Strategy    string
	PreferLocal bool
}

type CircuitBreakerConfig struct {
	Enabled         bool
	Threshold       float32
	WindowTime      int
	MinRequestCount int
	HalfOpenTime    time.Duration
	Check           CheckFunc
}

type BulkheadConfig struct {
	Enabled      bool
	Concurrency  int
	MaxQueueSize int
}

// Moleculer Config
type Config struct {
	// Namespace of nodes to segment your nodes on the same network (e.g.: “development”, “staging”, “production”). Default: ""
	Namespace string
	// Unique node identifier. Must be unique in a namespace. If not the broker will throw a fatal error and stop the process. Default: hostname + PID
	NodeID string
	// Logging configuration
	Logging LogConfig
	// Transit Config
	Transit TransitConfig

	// Transporter
	// Transporter string

	// // Request/Retry
	// MaxCallLevel   int
	// RequestTimeout time.Duration     // How long to wait before reject a request with a RequestTimeout error. Disabled: 0 Default: 0
	// RetryPolicy    RetryPolicyConfig // Retry policy configuration

	// // Heartbeat
	// HeartbeatTimeout  time.Duration
	// HeartbeatInterval time.Duration

	// Tracking TrackingConfig

	// DisableBalancer bool

	// Registry RegistryConfig

	// CircuitBreaker CircuitBreakerConfig

	// Bulkhead BulkheadConfig
}

func New(config Config) Config {

	return Config{
		Namespace: config.Namespace,
		NodeID:    getNodeID(config),
		Logging:   createDefaultLogConfig(config.Logging),
		Transit:   createDefaultTransitConfig(config.Transit),
	}
}

func getNodeID(config Config) string {
	if config.NodeID == "" {
		hostname, _ := os.Hostname()
		return fmt.Sprintf("%s%d", hostname, os.Getpid())
	}

	return config.NodeID
}
