// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"os"

	// homedir "github.com/mitchellh/go-homedir"
	"github.com/patrick-motard/dot/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// "fmt"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

var cfgFile string
var settings lib.Settings

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dot",
	Short: "A brief description of your application",
	Long:  `Dot is a CLI for interacting with your system.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	if os.Getenv("DEBUG") != "true" {
		log.SetLevel(logrus.WarnLevel)
	}
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dot.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// settings = lib.GetSettings()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)

	} else {
		// Find home directory.
		// home, err := homedir.Dir()
		// if err != nil {
		// 	log.Error(err)
		// 	os.Exit(1)
		// }

		// Search config in home directory with name ".dot" (without extension).

		var str = os.Getenv("GOPATH") + "/src/github.com/patrick-motard/dot/current_settings.toml"
		// viper.AddConfigPath("/home/han/code/go/src/github.com/patrick-motard/dot")
		// viper.SetConfigFile("/home/han/code/go/src/github.com/patrick-motard/dot/current_settings.toml")
		viper.SetConfigFile(str)
		// viper.SetConfigName("current_settings")
		// viper.SetConfigType("toml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Print("got here")
	// 	log.Info("Using config file:", viper.ConfigFileUsed())
	// }

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// for d, property := range viper.AllSettings() {
	// 	fmt.Println(property)
	// 	fmt.Println(d)
	// }

	// fmt.Println(len(viper.AllSettings()))
	// fmt.Println(viper.GetString("monitors.current"))
}

// example of setting a value and writing config:
// var newKeypair = make(map[string]string)
// newKeypair["test"] = "val"
// viper.Set("outerKey", newKeypair)
// viper.WriteConfig()
