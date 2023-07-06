package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

// Logger 全局日志记录器
var Logger *logrus.Logger

func NewLogger(level string) {
	Logger = logrus.New()
	Logger.SetFormatter(&customFormatter{})
	Logger.SetOutput(os.Stdout)

	switch strings.ToLower(level) {
	case "panic":
		Logger.SetLevel(logrus.PanicLevel)
		break
	case "fatal":
		Logger.SetLevel(logrus.FatalLevel)
		break
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
		break
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
		break
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
		break
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
		break
	case "trace":
		Logger.SetLevel(logrus.TraceLevel)
		break
	default:
		Logger.SetLevel(logrus.InfoLevel)
		break
	}
}

// 自定义日志格式化器
type customFormatter struct{}

func (customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	tagString := "main"
	if tag, exist := entry.Data["tag"]; exist {
		tagString = fmt.Sprint(tag)
	}

	var fields []string
	for key, value := range entry.Data {
		if key != "tag" {
			fields = append(fields, fmt.Sprintf("%s=%v", key, value))
		}
	}
	fieldsString := strings.Join(fields, " ")

	logString := fmt.Sprintf(
		"%s [%5s] %s: %s",
		entry.Time.Format("2006-01-02 15:04:05.000"),
		strings.ToUpper(entry.Level.String()),
		tagString,
		entry.Message,
	)
	if fieldsString != "" {
		logString += " { " + fieldsString + " }"
	}

	logString += "\n"

	return []byte(logString), nil
}
