package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

var addrCmd = &cobra.Command{
	Use:   "addr",
	Short: "Generate address information",
	Long: `
Generate Chinese address information`,
	Run: func(cmd *cobra.Command, args []string) {
		addr := generator.GetAddress()
		fmt.Println(addr)
		_ = clipboard.WriteAll(addr)
	},
}

func init() {
	rootCmd.AddCommand(addrCmd)
}
