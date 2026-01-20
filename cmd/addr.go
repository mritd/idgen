package cmd

import (
	"os"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid/v2"
	"github.com/mritd/idgen/v2/utils"
	"github.com/spf13/cobra"
)

var addrCmd = &cobra.Command{
	Use:   "addr",
	Short: "Generate address information",
	Long:  `Generate Chinese address information`,
	Run: func(cmd *cobra.Command, args []string) {
		var values []string
		for _, p := range chinaid.NewPerson().BuildN(count) {
			values = append(values, p.Address())
		}

		formatter := utils.NewFormatter(os.Stdout, format)
		_ = formatter.FormatSingle("address", values)

		if shouldCopy() {
			_ = clipboard.WriteAll(utils.SingleToClipboardText(values))
		}
	},
}

func init() {
	rootCmd.AddCommand(addrCmd)
}
