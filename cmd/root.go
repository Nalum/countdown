/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"math"
	"os"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile       string
	outputFile    string
	timerDuration time.Duration
	outputFormat  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "countdown",
	Short: "This tool is for use with the OBS text source.",
	Long: `Use this tool to create a count down timer, it will create a file and
update it as the timer counts down to 0.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := homedir.Expand(outputFile)

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		f, err := os.Create(path)
		defer f.Close()

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(2)
		}

		done := make(chan bool)
		ticker := time.NewTicker(time.Second)
		go startTimer(ticker.C, done, time.Duration(timerDuration), f)
		time.Sleep(timerDuration + time.Second)
		done <- true
		ticker.Stop()
	},
}

func startTimer(ticker <-chan time.Time, done <-chan bool, duration time.Duration, f *os.File) {
	for {
		select {
		case <-done:
			return
		case <-ticker:
			data := fmt.Sprintf("%02.0f:%02.0f", math.Floor(duration.Minutes()), duration.Seconds()-(math.Floor(duration.Minutes())*60))
			f.WriteAt([]byte(data), 0)
			duration = duration - time.Second

			if duration.Seconds() < 0 {
				duration = time.Duration(0)
			}
		}
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "config file (default is $HOME/.countdown.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&outputFile, "output-file", "o", "~/.countdown", "Where the output will be written")
	rootCmd.Flags().DurationVarP(&timerDuration, "duration", "d", time.Duration(5)*time.Minute, "Duration of the count down")
	rootCmd.Flags().StringVarP(&outputFormat, "format", "f", "mm:ss", "The format to output the count down timer in")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".countdown" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".countdown")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
