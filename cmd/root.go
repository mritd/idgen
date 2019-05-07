package cmd

import (
	"fmt"
	"os"

	"github.com/mritd/idgen/metadata"

	"github.com/atotto/clipboard"
	"github.com/mritd/idgen/generator"
	"github.com/spf13/cobra"
)

var version bool

var rootCmd = &cobra.Command{
	Use:   "idgen",
	Short: "身份信息生成器",
	Long: `
该工具用于生成中国大陆 姓名 身份证号 银行卡号 手机号 地址 Email
生成后自动复制相应文本到系统剪切板，不使用子命令则默认生成身份证号`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println("idgen:", metadata.VERSION)
		} else {
			idNo := generator.GetIDCard()
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
