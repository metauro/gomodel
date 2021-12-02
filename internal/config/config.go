package config

import (
	"github.com/metauro/gomodel/internal/msql"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	MySQL *msql.Config `toml:"mysql"`
}

func New(configFile string) (*Config, error) {
	viper.SetDefault("mysql", map[string]interface{}{
		"Username": "root",
		"Password": "wpsepmysql",
		"Type":     "mysql",
		"Host":     "localhost",
		"Port":     3306,
	})

	if configFile == "" {
		// Search config in current directory
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		viper.AddConfigPath(wd)

		// Search config in home directory
		home, err := homedir.Dir()
		if err != nil {
			return nil, err
		}
		viper.AddConfigPath(home)

		viper.SetConfigName(".gomodel")
	} else {
		viper.SetConfigFile(configFile)
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	log.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
