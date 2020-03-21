package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Coretabs-cli V 0.0.0",
	Long:  `Coretabs-cli V 0.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("V 0.0")
	},
}
