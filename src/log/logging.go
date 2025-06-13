package log

import (
	"Melodex/src/settings"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	var (
		logDir  = settings.AppConfig.LogDir
		logFile = settings.AppConfig.LogFile
		logPath = logDir + logFile
	)

	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		Log.Fatal("Failed to create log directory:", err)
	}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		Log.Fatal("Failed to open log file:", err)
	}

	Log.SetOutput(file)
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.InfoLevel)

	Log.Info("Log initialized")
}
