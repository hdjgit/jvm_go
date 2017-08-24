package log

import (
	"github.com/go-logging"
	"os"
)

var log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} [%{level:.5s}] %{id:03x}%{message}`,
)

func InitLog() {
	// For demo purposes, create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(os.Stdout, "", 0)
	//backend2 := logging.NewLogBackend(os.Stdout, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend1Formatter := logging.NewBackendFormatter(backend1, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Formatter)

}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}
