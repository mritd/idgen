package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

// addrCmd represents the addr command
var addrCmd = &cobra.Command{
	Use:   "addr",
	Short: "生成地址",
	Long: `
生成中国大陆地址信息`,
	Run: func(cmd *cobra.Command, args []string) {
		addr := generator.AddrGenerate()
		fmt.Println(addr)
		clipboard.WriteAll(addr)
	},
}

func init() {
	rootCmd.AddCommand(addrCmd)
}
