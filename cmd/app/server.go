package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teakingwang/gin-mysql/config"
	"github.com/teakingwang/gin-mysql/internal/app"
)

type Server struct {
}

var server = newServer()

func newServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	r := gin.Default()

	// Register routes
	userRoutes := r.Group("/api/v1/users")
	{
		userRoutes.GET("", app.GetUserList)
	}

	cfg := config.Config
	// Start server
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		logrus.Info("Failed to start server: %v", err)
	}

	select {}
}

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "server",
		Long:         `The server is the xeniro dataset management.`,
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			server.Run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringP("config", "c", "", "config file (default is $HOME/.cobra.yaml)")
	config.LoadConfig()

	viper.BindPFlags(cmd.Flags())

	return cmd
}
