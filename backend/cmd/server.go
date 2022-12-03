package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zoomoid/vox-populi/backend/cmd/options"
	"github.com/zoomoid/vox-populi/backend/pkg/config"
)

func NewServerCommand() *cobra.Command {

	serverConfig := options.NewServerConfiguration()

	cmd := &cobra.Command{
		Use:   "hob",
		Short: "Runs backend web server that manages polls",
		RunE: func(cmd *cobra.Command, args []string) error {
			var c *config.Configuration = serverConfig
			return newServer(c)
		},
	}

	options.AddFlags(serverConfig, cmd.Flags())

	return cmd
}

func newServer(conf *config.Configuration) error {
	server, cleanup, err := wireApp(conf)
	if err != nil {
		return err
	}
	defer cleanup()
	return server.Engine.Run()
}
