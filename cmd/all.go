package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate all information",
	Long: `
Generate all information`,
	Run: func(cmd *cobra.Command, args []string) {
		name := generator.GetName()
		idNo := generator.GetIDNo()
		mobile := generator.GetMobile()
		bank := generator.GetBank()
		email := generator.GetEmail()
		addr := generator.GetAddress()
		fmt.Println(name)
		fmt.Println(idNo)
		fmt.Println(mobile)
		fmt.Println(bank)
		fmt.Println(email)
		fmt.Println(addr)
		_ = clipboard.WriteAll(name + "\n" + idNo + "\n" + mobile + "\n" + bank + "\n" + email + "\n" + addr)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
