package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "生成姓名",
	Long: `
生成中文姓名，长度为 2-4 位(包含复姓)`,
	Run: func(cmd *cobra.Command, args []string) {
		name := generator.NameGenerate()
		fmt.Println(name)
		clipboard.WriteAll(name)
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)
}
