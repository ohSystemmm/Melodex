package settings

import (
	"os"
)

type Config struct {
	AppName    string
	Version    string
	Release    string
	SystemUser string
	ConfigDir  string
	CacheDir   string
	LogDir     string
	ConfigFile string
	CacheFile  string
	LogFile    string
}

var AppConfig = Config{
	AppName:    "melodex",
	Version:    "0.9.0",
	Release:    "07/01/2025",
	SystemUser: GetUser(),
	ConfigFile: "melodex.toml",
	LogFile:    "melodex.log",
}

func GetUser() string {
	envUser := os.Getenv("USER")
	if envUser == "" {
		envUser = "root"
	}
	return envUser
}

func init() {
	AppConfig.ConfigDir = "/home/" + AppConfig.SystemUser + "/.config/melodex/"
	AppConfig.CacheDir = "/home/" + AppConfig.SystemUser + "/.cache/melodex/cache/"
	AppConfig.LogDir = "/home/" + AppConfig.SystemUser + "/.cache/melodex/log/"
}
