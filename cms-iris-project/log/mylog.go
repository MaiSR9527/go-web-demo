package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

var Logger zerolog.Logger

func Log(level string, msg string) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	Log := zerolog.New(output).With().Timestamp().Logger()

	switch level {
	case "info":
		Log.Info().Msg(msg)
	case "error":
		Log.Error().Msg(msg)
	case "debug":
		Log.Debug().Msg(msg)
	case "":
		Log.Warn().Msg(msg)
	default:
		Log.Info().Msg(msg)
	}
}
