package cmd

import (
	"encoding/base64"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var bannerBase64 = `b29vb28gb29vb29vb29vby4gICAgICAgIC5vb29vb28uICAgIG9vb29vb29vb29vbyBvb29vbyAgICAgIG9vbyAKYDg4OCcgYDg4OCcgICBgWThiICAgICAgZDhQJyAgYFk4YiAgIGA4ODgnICAgICBgOCBgODg4Yi4gICAgIGA4JyAKIDg4OCAgIDg4OCAgICAgIDg4OCAgICA4ODggICAgICAgICAgICA4ODggICAgICAgICAgOCBgODhiLiAgICA4ICAKIDg4OCAgIDg4OCAgICAgIDg4OCAgICA4ODggICAgICAgICAgICA4ODhvb29vOCAgICAgOCAgIGA4OGIuICA4ICAKIDg4OCAgIDg4OCAgICAgIDg4OCAgICA4ODggICAgIG9vb29vICA4ODggICAgIiAgICAgOCAgICAgYDg4Yi44ICAKIDg4OCAgIDg4OCAgICAgZDg4JyAgICBgODguICAgIC44OCcgICA4ODggICAgICAgbyAgOCAgICAgICBgODg4ICAKbzg4OG8gbzg4OGJvb2Q4UCcgICAgICAgYFk4Ym9vZDhQJyAgIG84ODhvb29vb29kOCBvOG8gICAgICAgIGA4ICAK`

var versionTpl = `
%s
Name: idgen
Version: %s
Arch: %s
BuildDate: %s
CommitID: %s
`

var (
	Version   string
	BuildDate string
	CommitID  string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long: `
Print idgen version`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	banner, _ := base64.StdEncoding.DecodeString(bannerBase64)
	fmt.Printf(versionTpl, banner, Version, runtime.GOOS+"/"+runtime.GOARCH, BuildDate, CommitID)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
