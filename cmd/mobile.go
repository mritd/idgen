package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid"
	"github.com/spf13/cobra"
)

var mobileCmd = &cobra.Command{
	Use:   "mobile",
	Short: "Generate mobile phone number",
	Long: `
Generate mobile phone numbers in China`,
	Run: func(cmd *cobra.Command, args []string) {
		mobile := chinaid.Mobile()
		fmt.Println(mobile)
		_ = clipboard.WriteAll(mobile)
	},
}

func init() {
	rootCmd.AddCommand(mobileCmd)
}
