package initializers

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type CustomLogLevel zerolog.Level

const (
	InfoLevel  CustomLogLevel = CustomLogLevel(zerolog.InfoLevel)
	WarnLevel  CustomLogLevel = CustomLogLevel(zerolog.WarnLevel)
	ErrorLevel CustomLogLevel = CustomLogLevel(zerolog.ErrorLevel)
)

const (
	InfoLogFilePath  = "log/info.log"
	WarnLogFilePath  = "log/warn.log"
	ErrorLogFilePath = "log/error.log"
)

var logFiles struct {
	InfoLogFile  *os.File
	WarnLogFile  *os.File
	ErrorLogFile *os.File
}

type LoggerTypes struct {
	Info  *zerolog.Event
	Warn  *zerolog.Event
	Error *zerolog.Event
}

var Loggers LoggerTypes

func CustomZerologLogger() (infoLogger, warnLogger, errorLogger zerolog.Logger) {
	os.MkdirAll("log", os.ModePerm)

	infoFile := openFile(InfoLogFilePath)
	warnFile := openFile(WarnLogFilePath)
	errorFile := openFile(ErrorLogFilePath)

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("▶ %s", i)
		},
		FormatFieldName: func(i interface{}) string {
			return fmt.Sprintf("%s=", i)
		},
		FormatFieldValue: func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		},
		FormatCaller: func(i interface{}) string {
			return fmt.Sprintf(" %s", i)
		},
	}

	infoLogger = zerolog.New(io.MultiWriter(infoFile, consoleWriter)).
		Level(zerolog.InfoLevel).
		With().Timestamp().Caller().Logger()

	warnLogger = zerolog.New(io.MultiWriter(warnFile, consoleWriter)).
		Level(zerolog.WarnLevel).
		With().Timestamp().Caller().Logger()

	errorLogger = zerolog.New(io.MultiWriter(errorFile, consoleWriter)).
		Level(zerolog.ErrorLevel).
		With().Timestamp().Caller().Logger()

	return infoLogger, warnLogger, errorLogger
}
func init() {
	infoLogger, warnLogger, errorLogger := CustomZerologLogger()

	Loggers = LoggerTypes{
		Info:  infoLogger.Info(),
		Warn:  warnLogger.Warn(),
		Error: errorLogger.Error(),
	}

}

func LoggerCleanUp() {
	logFiles.InfoLogFile.Close()
	logFiles.WarnLogFile.Close()
	logFiles.ErrorLogFile.Close()
}

func openFile(LogFilePath string) *os.File {
	logFile, err := os.OpenFile(LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Failed to open log file: " + err.Error())
	}
	return logFile
}
