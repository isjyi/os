package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/isjyi/os/global"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

const defaultConfigFile = ".os.yaml"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "os",
	Short: "A test demo",
	Long:  `Demo is a test appcation for print things`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.os.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	v := viper.New()
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		v.SetConfigFile(defaultConfigFile)
	}

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(event fsnotify.Event) {
		fmt.Println("config file changed:", event.Name)
		if err := v.Unmarshal(&global.OS_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.OS_CONFIG); err != nil {
		fmt.Println(err)
	}

	global.OS_VP = v
}
