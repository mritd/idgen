package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

// mobileCmd represents the mobile command
var mobileCmd = &cobra.Command{
	Use:   "mobile",
	Short: "生成手机号",
	Long: `
生成大陆手机号`,
	Run: func(cmd *cobra.Command, args []string) {
		mobile := generator.MobileGenerate()
		fmt.Println(mobile)
		clipboard.WriteAll(mobile)
	},
}

func init() {
	rootCmd.AddCommand(mobileCmd)
}
