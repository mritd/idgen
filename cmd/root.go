package cmd

import (
	"os"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid/v2"
	"github.com/mritd/idgen/v2/utils"
	_ "github.com/mritd/logrus"
	"github.com/spf13/cobra"
)

// Global flags
var (
	count       int
	format      string
	doCopy      bool
	copyFlagSet bool
)

var rootCmd = &cobra.Command{
	Use:   "idgen",
	Short: "Identity information generator",
	Long: `Identity information generator for Chinese name, ID number,
bank card number, mobile phone number, address and Email.

Generate ID number by default without sub-command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default: generate ID number
		var values []string
		for _, p := range chinaid.NewPerson().BuildN(count) {
			values = append(values, p.IDNo())
		}

		formatter := utils.NewFormatter(os.Stdout, format)
		_ = formatter.FormatSingle("idno", values)

		if shouldCopy() {
			_ = clipboard.WriteAll(utils.SingleToClipboardText(values))
		}
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&count, "count", "c", 1, "Number of records to generate")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "table", "Output format: table|json|csv")
	rootCmd.PersistentFlags().BoolVarP(&doCopy, "copy", "C", false, "Copy to clipboard (default: true for single, false for batch)")

	// Track if copy flag was explicitly set
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		copyFlagSet = cmd.Flags().Changed("copy")
	}
}

// shouldCopy returns whether to copy to clipboard
func shouldCopy() bool {
	// If explicitly set, use the flag value
	if copyFlagSet {
		return doCopy
	}
	// Default: copy for single record, don't copy for batch
	return count == 1
}
