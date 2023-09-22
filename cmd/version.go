package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version",
	Long:  `Show the current version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(AppName, AppVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
