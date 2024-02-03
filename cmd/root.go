/*
Copyright © 2024 Rashi M rashi1281@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rashiraffi/trai/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "trai",
	Short: "A CLI tool bringing AI to the terminal, enabling users to interact with artificial intelligence seamlessly.",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//
	//	Run: func(cmd *cobra.Command, args []string) {
	//		if len(args) != 0 {
	//			ask.Ask(context.Background(), strings.Join(args, " "))
	//		}
	//	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.trai.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".trai" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".trai")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file not found, setting up configuration")
		err := config.SetConfig()
		cobra.CheckErr(err)
		viper.SafeWriteConfig()
	}
}
