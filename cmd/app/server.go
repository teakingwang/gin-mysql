package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teakingwang/gin-mysql/config"
	"github.com/teakingwang/gin-mysql/pkg/db"
	"net"
)

type Server struct {
	router *Router
}

var server = newServer()

func newServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	// load config
	config.LoadConfig()
	// 初始化db
	db.InitDB()
	// 数据库迁移
	db.MigrateDB(db.GormDB)

	// router
	s.router = NewRouter(net.JoinHostPort(config.Config.Server.Host, config.Config.Server.Port))
	s.router.Config()
	s.router.Run()

	select {}
}

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "server",
		Long:         `The server is demo mvc project`,
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
	viper.BindPFlags(cmd.Flags())

	return cmd
}
