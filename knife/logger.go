package knife

import (
	"log"
	"os"
	"time"
)

var Logger = log.New(os.Stdout,
	time.Now().Format("2006-01-02 15:04:05.000 "),
	log.Lshortfile,
)
