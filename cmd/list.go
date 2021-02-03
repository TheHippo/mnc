package cmd

import (
	"fmt"

	"github.com/TheHippo/mnc/device"
	"github.com/TheHippo/mnc/util"

	"github.com/spf13/cobra"
)

var listOnlyIP4 bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists available network interfaces",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ifaces, err := device.GetInterfaces(listOnlyIP4)
		if err != nil {
			return err
		}
		fmt.Print(util.MakePretty(device.MakePrettyIter(ifaces)))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&listOnlyIP4, "only-ip4", "4", true, "list only IP4 adresses")

}
