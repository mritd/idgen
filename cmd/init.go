package cmd

import (
	"github.com/mritd/idgen/utils"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init config",
	Long: `
Init idgen config(eg: bbolt db file)`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ExportDB()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
