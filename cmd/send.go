package cmd

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TheHippo/mnc/transfer"
	"github.com/TheHippo/mnc/util"
	"github.com/spf13/cobra"
)

var calculateSendHash bool

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send stdin to other interfaces",
	RunE: func(c *cobra.Command, args []string) error {
		if !util.IsPipe(os.Stdin) {
			logger.Fatalln("No data in stdin")
		}
		md5 := transfer.NewMD5Reader(os.Stdin, calculateSendHash)
		_, err := ioutil.ReadAll(md5)
		if err != nil {
			fmt.Println(err)
		}
		if md5.Hashing() {
			fmt.Printf("%s\n", hex.EncodeToString(md5.Sum()))
		}
		return nil
	},
}

func init() {
	sendCmd.Flags().BoolVarP(&calculateSendHash, "hashing", "m", true, "check transmitted data via md5 hash")
	rootCmd.AddCommand(sendCmd)
}
