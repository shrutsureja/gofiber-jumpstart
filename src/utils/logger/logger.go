package logger

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Customize the caller field to just show the filename and line number
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
    	return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	return log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
}