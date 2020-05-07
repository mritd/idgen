package cmd

import (
	"fmt"

	"github.com/mritd/chinaid"

	"github.com/atotto/clipboard"

	"github.com/spf13/cobra"
)

var bankCmd = &cobra.Command{
	Use:   "bank",
	Short: "Generate bank card number",
	Long: `
Generate Bank of China Debit Card Number`,
	Run: func(cmd *cobra.Command, args []string) {
		bank := chinaid.BankNo()
		fmt.Println(bank)
		_ = clipboard.WriteAll(bank)
	},
}

func init() {
	rootCmd.AddCommand(bankCmd)
}
