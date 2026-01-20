package cmd

import (
	"os"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid/v2"
	"github.com/mritd/idgen/v2/utils"
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate all information",
	Long:  `Generate all information including name, ID, mobile, bank, email and address`,
	Run: func(cmd *cobra.Command, args []string) {
		var identities []utils.Identity
		for _, p := range chinaid.NewPerson().BuildN(count) {
			identities = append(identities, utils.Identity{
				Name:    p.Name(),
				IDNo:    p.IDNo(),
				Mobile:  p.Mobile(),
				Bank:    p.BankNo(),
				Email:   p.Email(),
				Address: p.Address(),
			})
		}

		formatter := utils.NewFormatter(os.Stdout, format)
		_ = formatter.FormatIdentities(identities)

		if shouldCopy() {
			_ = clipboard.WriteAll(utils.ToClipboardText(identities))
		}
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
