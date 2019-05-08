package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "Generate email address",
	Long: `
Generate Email in the format of "8 lowercase letters"@"5 lowercase letters"."Common top level domain suffixes"`,
	Run: func(cmd *cobra.Command, args []string) {
		email := generator.GetEmail()
		fmt.Println(email)
		_ = clipboard.WriteAll(email)
	},
}

func init() {
	rootCmd.AddCommand(emailCmd)
}
