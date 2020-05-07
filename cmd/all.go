package cmd

import (
	"fmt"

	"github.com/mritd/chinaid"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate all information",
	Long: `
Generate all information`,
	Run: func(cmd *cobra.Command, args []string) {
		name := chinaid.Name()
		idNo := chinaid.IDNo()
		mobile := chinaid.Mobile()
		bank := chinaid.BankNo()
		email := chinaid.Email()
		addr := chinaid.Address()
		fmt.Println(name)
		fmt.Println(idNo)
		fmt.Println(mobile)
		fmt.Println(bank)
		fmt.Println(email)
		fmt.Println(addr)
		_ = clipboard.WriteAll(name + "\n" + idNo + "\n" + mobile + "\n" + bank + "\n" + email + "\n" + addr)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
