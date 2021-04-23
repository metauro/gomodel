/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/jmoiron/sqlx"
	"github.com/metauro/gomodel/internal/msql"
	"github.com/spf13/cobra"
)

var cfgFile string
var db *sqlx.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomodel",
	Short: "模型代码生成器",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// Find home directory.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "配置文件 (default is $HOME/.gomodel.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config, err := NewConfig()
	cobra.CheckErr(err)
	config.MySQL.Type = "mysql"
	db = msql.Open(config.MySQL)
}
