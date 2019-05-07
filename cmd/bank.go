package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"

	"github.com/spf13/cobra"
)

// bankCmd represents the bank command
var bankCmd = &cobra.Command{
	Use:   "bank",
	Short: "生成银行卡号",
	Long: `
生成中国大部分银行的银行卡号`,
	Run: func(cmd *cobra.Command, args []string) {
		bank := generator.BankGenerate()
		fmt.Println(bank)
		clipboard.WriteAll(bank)
	},
}

func init() {
	rootCmd.AddCommand(bankCmd)
}
