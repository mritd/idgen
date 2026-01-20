package cmd

import (
	"github.com/mritd/idgen/v2/server"
	"github.com/spf13/cobra"
)

var (
	listen string
	port   int
	theme  string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run as http server",
	Long: `Run a simple http server to provide page access and API.

Access endpoints:
  http://BINDADDR:PORT/              HTML page with theme
  http://BINDADDR:PORT/api/v1/generate   Generate single record (JSON)
  http://BINDADDR:PORT/api/v1/batch      Batch generate (JSON)
  http://BINDADDR:PORT/api/v1/export     Export as CSV`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start(listen, port, theme)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&listen, "listen", "l", "0.0.0.0", "HTTP listen address")
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "HTTP listen port")
	serverCmd.Flags().StringVarP(&theme, "theme", "t", "cyber", "Default theme: cyber|terminal")
}
