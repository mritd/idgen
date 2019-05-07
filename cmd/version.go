package cmd

import (
	"fmt"

	"github.com/mritd/idgen/metadata"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示当前版本",
	Long: `
显示当前版本`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("idgen:", metadata.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
