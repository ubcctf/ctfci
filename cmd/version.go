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
	Short: "Print the version",
	Long:  `Print the version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ctfci-v0.0.0-dev")
	},
}
