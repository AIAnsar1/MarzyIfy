package Logger

import (
	"github/rs/zerolog"
)

type NoopLogger struct {
}

type ContextFilterHook struct {
	ContextKey   string
	ContextValue string
}

const (
	LOG_CONTEXT = "log-context"
)

var (
	Logger zerolog.Logger
)

func init() {

}

func (this *ContextFilterHook) MarzyLogRun(E *zerolog.Eventm, Level zerolog.Level, Message string) {

}
