package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

// idnoCmd represents the idno command
var idnoCmd = &cobra.Command{
	Use:   "idno",
	Short: "生成身份证号",
	Long: `
生成中国大陆十八位身份证号`,
	Run: func(cmd *cobra.Command, args []string) {
		idNo := generator.IDCardGenerate()
		fmt.Println(idNo)
		clipboard.WriteAll(idNo)
	},
}

func init() {
	rootCmd.AddCommand(idnoCmd)
}
