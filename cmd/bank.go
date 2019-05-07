package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"

	"github.com/spf13/cobra"
)

var bankCmd = &cobra.Command{
	Use:   "bank",
	Short: "生成银行卡号",
	Long: `
生成中国大部分银行的银行卡号`,
	Run: func(cmd *cobra.Command, args []string) {
		bank := generator.GetBank()
		fmt.Println(bank)
		_ = clipboard.WriteAll(bank)
	},
}

func init() {
	rootCmd.AddCommand(bankCmd)
}
