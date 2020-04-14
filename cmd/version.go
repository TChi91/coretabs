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
	Short: "Print version information and quit",
	Long:  `Print version information and quit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("V 0.0")
	},
}
