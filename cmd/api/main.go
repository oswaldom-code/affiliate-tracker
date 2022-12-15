package api

import (
	"os"

	"github.com/oswaldom-code/affiliate-tracker/pkg/config"
	"github.com/oswaldom-code/affiliate-tracker/src/adapters/http/rest/infrastructure"
	"github.com/spf13/cobra"
)

// load config file and set environment variables
func init() {
	config.LoadConfigurationFile()
	config.SetEnvironment(config.GetEnvironmentConfig().Environment)
}

// serveCmd represents the serve command
var serveCmdNew = &cobra.Command{
	Use:   "server",
	Short: "Spin up the web server that hosts the API",
	Long:  `The web server hosts the API, and manages the authentication middleware`,
	Run: func(cmd *cobra.Command, args []string) {
		NewServer()
	},
}

func NewServer() {
	r := infrastructure.NewServer()
	var err error
	//setup routes
	println("Server running at: ", config.GetServerConfig().AsUri())
	if config.GetServerConfig().Scheme == "https" { // https
		err = r.RunTLS(
			config.GetServerConfig().AsUri(),
			config.GetServerConfig().PathToSSLCertFile,
			config.GetServerConfig().PathToSSLKeyFile,
		)
	} else { // http
		err = r.Run(config.GetServerConfig().AsUri())
	}
	if err != nil {
		panic(err.Error())
	}
}

// Execute runs the root command
func Execute() {
	if err := serveCmdNew.Execute(); err != nil {
		os.Exit(1)
	}
}
