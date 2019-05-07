package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "生成所有信息",
	Long: `
生成本工具支持的所有信息`,
	Run: func(cmd *cobra.Command, args []string) {
		name := generator.NameGenerate()
		idNo := generator.IDCardGenerate()
		mobile := generator.MobileGenerate()
		bank := generator.BankGenerate()
		email := generator.EmailGenerate()
		addr := generator.AddrGenerate()
		fmt.Println(name)
		fmt.Println(idNo)
		fmt.Println(mobile)
		fmt.Println(bank)
		fmt.Println(email)
		fmt.Println(addr)
		clipboard.WriteAll(name + "\n" + idNo + "\n" + mobile + "\n" + bank + "\n" + email + "\n" + addr)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
