package logging

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var Log *logrus.Logger

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()

	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}

	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func New() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	currentTime := time.Now()
	fileName := currentTime.Format("log_2006-01-02.log")

	os.Mkdir("logs", 0644)

	file, err := os.OpenFile("logs/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Info("Error opening file, by default will use a output to a console")
	}

	log.AddHook(&writerHook{
		LogLevels: logrus.AllLevels,
		Writer:    []io.Writer{file},
	})

	log.SetLevel(logrus.InfoLevel)

	return log
}

func Init() {
	Log = New()
}
