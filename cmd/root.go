package cmd

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

var version bool

var rootCmd = &cobra.Command{
	Use:   "idgen",
	Short: "Identity information generator",
	Long: `
This tool is used to generate Chinese name、ID number、bank card number、
mobile phone number、address and Email; automatically generate corresponding
text to the system clipboard after generation, and generate ID number by
default without sub-command`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			printVersion()
		} else {
			idNo := generator.GetIDNo()
			fmt.Println(idNo)
			_ = clipboard.WriteAll(idNo)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "显示当前版本")
}
