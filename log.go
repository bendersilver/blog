package blog

import (
	"fmt"
	"os"
	"path"

	"github.com/op/go-logging"
)

var logger *logging.Logger

var formatSys = logging.MustStringFormatter(
	`%{color}%{level:.1s} %{time:2006-01-02 15:04:05.000} %{shortpkg} %{shortfile} ▶ %{color:reset} %{message}`,
)

// init -
func init() {
	logger = logging.MustGetLogger(path.Base(os.Args[0]))
	logger.ExtraCalldepth = 1
	logging.SetBackend(
		logging.NewBackendFormatter(
			logging.NewLogBackend(os.Stderr, "", 0), formatSys,
		),
	)
}

// Debug -
func Debug(v ...interface{}) {
	logger.Debug(fmt.Sprintln(v...))
}

// Debugf -
func Debugf(format string, a ...interface{}) {
	logger.Debug(fmt.Sprintf(format, a...))
}

// Info -
func Info(v ...interface{}) {
	logger.Info(fmt.Sprintln(v...))
}

// Infof -
func Infof(format string, a ...interface{}) {
	logger.Info(fmt.Sprintf(format, a...))
}

// Notice -
func Notice(v ...interface{}) {
	logger.Notice(fmt.Sprintln(v...))
}

// Noticef -
func Noticef(format string, a ...interface{}) {
	logger.Notice(fmt.Sprintf(format, a...))
}

// Warning -
func Warning(v ...interface{}) {
	logger.Warning(fmt.Sprintln(v...))
}

// Warningf -
func Warningf(format string, a ...interface{}) {
	logger.Warning(fmt.Sprintf(format, a...))
}

// Error -
func Error(v ...interface{}) {
	logger.Error(fmt.Sprintln(v...))
}

// Errorf -
func Errorf(format string, a ...interface{}) {
	logger.Error(fmt.Sprintf(format, a...))
}

// Fatal -
func Fatal(v ...interface{}) {
	logger.Critical(fmt.Sprintln(v...))
	os.Exit(1)
}

// Fatalf -
func Fatalf(format string, a ...interface{}) {
	logger.Critical(fmt.Sprintf(format, a...))
	os.Exit(1)
}
