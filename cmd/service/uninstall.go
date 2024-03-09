package service

import (
	"fmt"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

func NewSvcUnInstallCmd(s service.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "卸载服务",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("卸载服务")
			err := service.Control(s, "uninstall")
			if err != nil {
				logx.Must(err)
			}
		},
	}
}
