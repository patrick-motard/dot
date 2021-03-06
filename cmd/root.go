// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"os"
	"os/user"

	"github.com/patrick-motard/dot/lib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Home string

var log = logrus.New()

// TODO: make this path a public variable
var cfgFile string
var settings lib.Settings

type Theme struct {
	Name string
	Bars []string
	Gaps I3Gaps
}

type I3Gaps struct {
	Top    string
	Bottom string
	Left   string
	Right  string
}
type I3wm struct {
	DefaultGaps  I3Gaps `mapstructure:"default_gaps"`
	SettingsFile string `mapstructure:"settings_file"`
}

type config struct {
	Displays struct {
		Current  string
		Location string
	}
	Sound struct {
		Port string
	}
	I3wm    I3wm `mapstructure:"i3_wm"`
	Polybar struct {
		Theme           string
		ThemesDirectory string `mapstructure:"themes_directory"`
		Themes          []Theme
	}
}

var Config config

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
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	Home = usr.HomeDir
	if os.Getenv("DEBUG") != "true" {
		log.SetLevel(logrus.InfoLevel)
	}
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dot.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)

	} else {
		cfgFile = Home + "/code/dot/current_settings.yml"
		viper.SetConfigFile(cfgFile)
	}

	viper.AutomaticEnv() // read in environment variables that match
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("Failed to parse config file %s\nError Message: %s", cfgFile, err)
		os.Exit(1)
	}
	uErr := viper.Unmarshal(&Config)
	if uErr != nil {
		log.Fatalf("Unable to decode config into struct, %v", uErr)
	}
}

// example of setting a value and writing config:
// var newKeypair = make(map[string]string)
// newKeypair["test"] = "val"
// viper.Set("outerKey", newKeypair)
// viper.WriteConfig()
