package service_broker

// import (
// 	"github.com/olebedev/emitter"
// 	"github.com/rackworx/moleculer-go/pkg/config"
// 	sb "github.com/rackworx/moleculer-go/pkg/service_broker"
// 	"github.com/rs/zerolog"
// 	"github.com/rs/zerolog/log"
// )

// type serviceBroker struct {
// 	config config.Config
// 	bus    *emitter.Emitter
// 	logger zerolog.Logger
// }

// func New(config config.Config) sb.ServiceBroker {
// 	logger := log.Output(zerolog.ConsoleWriter{Out: config.Logging.Out})

// 	return &serviceBroker{
// 		config: config,
// 		bus:    &emitter.Emitter{},
// 		logger: logger.With().Str("component", "broker").Logger(),
// 	}
// }

// func (s *serviceBroker) Start() {
// 	// starTime := time.Now()
// }

// func (s *serviceBroker) GetLogger(name string) zerolog.Logger {
// 	return s.logger.With().Str("component", name).Logger()
// }
