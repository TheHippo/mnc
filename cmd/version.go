package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "undefined"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows the version of mnc you are using",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
