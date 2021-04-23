package cmd

import (
	"github.com/metauro/gomodel/internal/msql"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	MySQL *msql.Config `toml:"mysql"`
}

func NewConfig() (*Config, error) {
	viper.SetDefault("mysql", map[string]interface{}{
		"Username": "root",
		"Password": "root",
		"Type":     "mysql",
		"Host":     "localhost",
		"Port":     3306,
	})

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return nil, err
		}

		// Search config in home directory with name ".gomodel" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gomodel")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	log.Printf("Using config file: %s", viper.ConfigFileUsed())
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
