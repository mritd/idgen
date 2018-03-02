// Copyright © 2018 mritd <mritd1234@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/mritd/idgen/metadata"
	"os"

	"github.com/atotto/clipboard"
	"github.com/mitchellh/go-homedir"
	"github.com/mritd/idgen/generator"
	"github.com/mritd/idgen/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
)

var cfgFile string

var version bool

// rootCmd represents the base command when called without any subcommands
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
			idNo := generator.IDCardGenerate()
			fmt.Println(idNo)
			clipboard.WriteAll(idNo)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.idgen/idgen.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "显示当前版本")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		util.CheckAndExit(err)

		// Search config in home directory with name ".idgen" (without extension).
		viper.AddConfigPath(home + string(filepath.Separator) + ".idgen")
		viper.SetConfigName("idgen")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
