package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "生成姓名",
	Long: `
生成中文姓名，长度为 2-4 位(包含复姓)`,
	Run: func(cmd *cobra.Command, args []string) {
		name := generator.GetName()
		fmt.Println(name)
		_ = clipboard.WriteAll(name)
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)
}
