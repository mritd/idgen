package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"

	"github.com/spf13/cobra"
)

var bankCmd = &cobra.Command{
	Use:   "bank",
	Short: "Generate bank card number",
	Long: `
Generate Bank of China Debit Card Number`,
	Run: func(cmd *cobra.Command, args []string) {
		bank := generator.GetBank()
		fmt.Println(bank)
		_ = clipboard.WriteAll(bank)
	},
}

func init() {
	rootCmd.AddCommand(bankCmd)
}
