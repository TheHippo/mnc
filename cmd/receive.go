package cmd

import "github.com/spf13/cobra"

var calculateReceiveHash bool
var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "receive data from somewhere",
	RunE: func(c *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	receiveCmd.Flags().BoolVarP(&calculateReceiveHash, "hashing", "m", true, "check transmitted data via md5 hash")
	rootCmd.AddCommand(receiveCmd)
}
