package cmd

import (
	"net"
	"strconv"

	"github.com/mritd/idgen/server"
	"github.com/mritd/idgen/utils"
	"github.com/spf13/cobra"
)

var listen *string
var port *int
var mode *string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动 http server",
	Long: `
启动一个简单的 http server 用于提供页面访问以及 json 数据返回，
当不指定 -m 选项则同时开启 html 和 json 支持，访问地址如下:

http://BINDADDR:PORT/        返回一个简单的 html 页面
http://BINDADDR:PORT/api     返回 json 格式数据`,
	Run: func(cmd *cobra.Command, args []string) {
		tcpAddr, err := net.ResolveTCPAddr("tcp", *listen+":"+strconv.Itoa(*port))
		utils.CheckAndExit(err)
		server.Start(*mode, tcpAddr)
	},
}

func init() {

	rootCmd.AddCommand(serverCmd)
	mode = serverCmd.PersistentFlags().StringP("mode", "m", "", "server 运行模式(html/json)")
	listen = serverCmd.PersistentFlags().StringP("listen", "l", "0.0.0.0", "http 监听地址")
	port = serverCmd.PersistentFlags().IntP("port", "p", 8080, "http 监听端口")
}
