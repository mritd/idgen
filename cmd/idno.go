package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid"
	"github.com/spf13/cobra"
)

var idnoCmd = &cobra.Command{
	Use:   "idno",
	Short: "Generate ID number",
	Long: `
Generate 18-digit ID number of Chinese citizens`,
	Run: func(cmd *cobra.Command, args []string) {
		idNo := chinaid.IDNo()
		fmt.Println(idNo)
		_ = clipboard.WriteAll(idNo)
	},
}

func init() {
	rootCmd.AddCommand(idnoCmd)
}
