package cmd

import (
	"github.com/mritd/idgen/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化配置",
	Long: `
初始化基本配置，如从网络下载用于生成姓名的 SQLite 数据库等`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.InitDB()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
