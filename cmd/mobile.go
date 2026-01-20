package cmd

import (
	"os"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid/v2"
	"github.com/mritd/idgen/v2/utils"
	"github.com/spf13/cobra"
)

var mobileCmd = &cobra.Command{
	Use:   "mobile",
	Short: "Generate mobile phone number",
	Long:  `Generate mobile phone numbers in China`,
	Run: func(cmd *cobra.Command, args []string) {
		var values []string
		for _, p := range chinaid.NewPerson().BuildN(count) {
			values = append(values, p.Mobile())
		}

		formatter := utils.NewFormatter(os.Stdout, format)
		_ = formatter.FormatSingle("mobile", values)

		if shouldCopy() {
			_ = clipboard.WriteAll(utils.SingleToClipboardText(values))
		}
	},
}

func init() {
	rootCmd.AddCommand(mobileCmd)
}
