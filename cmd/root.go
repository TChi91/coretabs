package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// This represents the base command when called
	// without any subcommands

	rootCmd = &cobra.Command{
		Use:   "coretabs",
		Short: "We generate your project ",
		Long: `Coretabs-cli will generate front-end, back-end projects, and both for you.
You just have to focus on code :p`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		er(err)
	}

	// Search config in home directory with name ".cobra" (without extension).
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(home + "/.coretabs/")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Println(err)
	// }
}
