package cmd

import (
	"os"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid/v2"
	"github.com/mritd/idgen/utils"
	"github.com/spf13/cobra"
)

var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "Generate email address",
	Long:  `Generate Email in the format of "8 lowercase letters"@"5 lowercase letters"."Common TLD"`,
	Run: func(cmd *cobra.Command, args []string) {
		var values []string
		for _, p := range chinaid.NewPerson().BuildN(count) {
			values = append(values, p.Email())
		}

		formatter := utils.NewFormatter(os.Stdout, format)
		_ = formatter.FormatSingle("email", values)

		if shouldCopy() {
			_ = clipboard.WriteAll(utils.SingleToClipboardText(values))
		}
	},
}

func init() {
	rootCmd.AddCommand(emailCmd)
}
