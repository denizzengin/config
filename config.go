package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	configPath          string = "./"
	configFileType      string = "yaml"
	configName          string = "config"
	configFileExtension string = ".yml"
	productionEnv       string = "Production"
)

var Config *Configuration = New()

func New() (*Configuration) {

	viper.SetConfigType(configFileType)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	viper.BindEnv("environment.env", "B_ENV")
	viper.SetDefault("environment.env", "Development")

	configFilePath := filepath.Join(configPath, configName) + configFileExtension
	if err := readConfiguration(configFilePath); err != nil {
		return nil
	}

	viper.AutomaticEnv()
	var cfg *Configuration
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		if cfg.Environment.Env != productionEnv {
			log.Println("Config file changed:", e.Name)
		}
	})

	viper.WatchConfig()

	return cfg
}

// read configuration from file
func readConfiguration(filePath string) error {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		// if file does not exist, simply create one
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			os.Create(filePath)
		} else {
			return err
		}
		// let's write defaults
		if err := viper.WriteConfig(); err != nil {
			return err
		}
	}
	return nil
}
