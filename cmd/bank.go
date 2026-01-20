package cmd

import (
	"os"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid/v2"
	"github.com/mritd/idgen/v2/utils"
	"github.com/spf13/cobra"
)

var bankCmd = &cobra.Command{
	Use:   "bank",
	Short: "Generate bank card number",
	Long:  `Generate Bank of China Debit Card Number`,
	Run: func(cmd *cobra.Command, args []string) {
		var values []string
		for _, p := range chinaid.NewPerson().BuildN(count) {
			values = append(values, p.BankNo())
		}

		formatter := utils.NewFormatter(os.Stdout, format)
		_ = formatter.FormatSingle("bank", values)

		if shouldCopy() {
			_ = clipboard.WriteAll(utils.SingleToClipboardText(values))
		}
	},
}

func init() {
	rootCmd.AddCommand(bankCmd)
}
