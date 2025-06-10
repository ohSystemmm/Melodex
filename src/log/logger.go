package log

import (
	"Melodex/src/settings"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	err := os.MkdirAll(settings.AppConfig.LogDir, 0755)
	if err != nil {
		Log.Fatal("Failed to create log directory:", err)
	}

	logFilePath := filepath.Join(settings.AppConfig.LogDir, settings.AppConfig.LogFile)
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		Log.Fatal("Failed to open log file:", err)
	}

	Log.SetOutput(file)
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.InfoLevel)

	Log.Info("Log initialized")
}
