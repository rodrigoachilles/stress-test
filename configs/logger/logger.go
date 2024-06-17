package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var (
	log zerolog.Logger
)

func init() {
	log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()

}

func Info(message string) {
	log.Info().Msg(message)
}
