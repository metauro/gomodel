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
	Use:     "gomodel",
	Short:   "model code generator",
	Example: "gomodel gen",
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", `Config file path (default "$HOME/.gomodel.toml")`)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config, err := NewConfig()
	cobra.CheckErr(err)
	db = msql.Open(config.MySQL)
}
