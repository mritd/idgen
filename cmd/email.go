package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

// emailCmd represents the email command
var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "生成 Email",
	Long: `
生成 Email，格式为 8位小写字母@5位小写字母.常用顶级域名后缀`,
	Run: func(cmd *cobra.Command, args []string) {
		email := generator.EmailGenerate()
		fmt.Println(email)
		clipboard.WriteAll(email)
	},
}

func init() {
	rootCmd.AddCommand(emailCmd)
}
