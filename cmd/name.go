package cmd

import (
	"os"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid/v2"
	"github.com/mritd/idgen/v2/utils"
	"github.com/spf13/cobra"
)

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Generate name",
	Long:  `Generate Chinese name, length 2-4 digits (including complex surname)`,
	Run: func(cmd *cobra.Command, args []string) {
		var values []string
		for _, p := range chinaid.NewPerson().BuildN(count) {
			values = append(values, p.Name())
		}

		formatter := utils.NewFormatter(os.Stdout, format)
		_ = formatter.FormatSingle("name", values)

		if shouldCopy() {
			_ = clipboard.WriteAll(utils.SingleToClipboardText(values))
		}
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)
}
