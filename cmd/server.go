package cmd

import (
	"net"
	"strconv"

	"github.com/mritd/idgen/server"
	"github.com/spf13/cobra"
)

var listen string
var port int
var mode string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run as http server",
	Long: `
Run a simple http server to provide page access and json data return.
When the -m option is not specified, both html and json support are enabled, 
and the access address is as follows:

http://BINDADDR:PORT/        return a simple html page
http://BINDADDR:PORT/api     return json format data`,
	Run: func(cmd *cobra.Command, args []string) {
		tcpAddr, err := net.ResolveTCPAddr("tcp", listen+":"+strconv.Itoa(port))
		if err != nil {
			panic(err)
		}
		server.Start(mode, tcpAddr)
	},
}

func init() {

	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "", "server mode(html/json)")
	serverCmd.PersistentFlags().StringVarP(&listen, "listen", "l", "0.0.0.0", "http listen address")
	serverCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "http listen port")
}
