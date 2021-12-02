package cmd

import (
	"github.com/metauro/gomodel/internal/msql"
	"github.com/spf13/cobra"
)

var driveName string
var dsn string
var db *msql.DB

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
	rootCmd.PersistentFlags().StringVar(&dsn, "dsn", "", "eg: root:root@(localhost:3306)/test?parseTime=true")
	rootCmd.PersistentFlags().StringVar(&driveName, "drive-name", "mysql", "mysql,postgres")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var err error
	db, err = msql.Open(driveName, dsn)
	if err != nil {
		panic(err)
	}
}
