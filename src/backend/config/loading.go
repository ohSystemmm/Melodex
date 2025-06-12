package config

//func LoadConfig(cfgFile string) (Config, error) {
//	if cfgFile == "" {
//		cfgFile = filepath.Join(settings.AppConfig.ConfigDir, settings.AppConfig.ConfigFile)
//	}
//
//	data, err := os.ReadFile(cfgFile)
//	if err != nil {
//		log.Log.Warnf("Config file not found: %s. Falling back to default.", cfgFile)
//		return GenerateDefaultConfig(), err
//	}
//
//	var cfg Config
//	if err = toml.Unmarshal(data, &cfg); err != nil {
//		log.Log.Warnf("Error parsing config file: %s. Using default settings.", cfgFile)
//		return GenerateDefaultConfig(), err
//	}
//
//	return cfg, nil
//}
