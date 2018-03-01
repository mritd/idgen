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
	"github.com/mritd/idgen/server"
	"github.com/mritd/idgen/util"
	"github.com/spf13/cobra"
	"net"
	"strconv"
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
		util.CheckAndExit(err)
		server.Start(*mode, tcpAddr)
	},
}

func init() {

	rootCmd.AddCommand(serverCmd)
	mode = serverCmd.PersistentFlags().StringP("mode", "m", "", "server 运行模式(html/json)")
	listen = serverCmd.PersistentFlags().StringP("listen", "l", "0.0.0.0", "http 监听地址")
	port = serverCmd.PersistentFlags().IntP("port", "p", 8080, "http 监听端口")
}
