package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func InitLogrus() {
	// Set log format as json
	Log.Formatter = &logrus.JSONFormatter{}
	file, err := os.OpenFile("./gin_log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("Failed to create/open log file: %s \n", err))
	}
	// Set log default output file
	Log.Out = file
	// Set log level
	Log.Level = logrus.DebugLevel
}
