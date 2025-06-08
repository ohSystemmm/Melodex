package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var Log *logrus.Logger

// init initializes the logging system by creating necessary directories and setting up log output.
// The log file is stored in the user's cache directory under "melodex/app.log".
func init() {
	Log = logrus.New()

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		Log.Fatal("Failed to get user cache directory:", err)
	}

	logDir := filepath.Join(cacheDir, "melodex")
	logFilePath := filepath.Join(logDir, "app.log")

	err = os.MkdirAll(logDir, 0755)
	if err != nil {
		Log.Fatal("Failed to create log directory:", err)
	}

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatal("Failed to open log file:", err)
	}

	Log.SetOutput(file)
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.InfoLevel)
}
