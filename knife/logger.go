package knife

import (
	"log"
	"os"
	"time"
)

var Logger = newLogger()

type loggers struct {
	*log.Logger
}

func newLogger() *loggers {
	now := time.Now().Format("2006-01-02 15:04:05.000")
	l := log.New(os.Stdout, now+" ", log.Lshortfile)
	return &loggers{
		Logger: l,
	}
}
