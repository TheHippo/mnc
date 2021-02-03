package cmd

import (
	"github.com/spf13/cobra"
)

var version = "undefined"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows the version of mnc you are using",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
