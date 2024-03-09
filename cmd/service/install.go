package service

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

func NewSvcInstallCmd(s service.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "安装服务",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("安装服务")
			err := service.Control(s, "install")
			if err != nil {
				logx.Must(err)
			}
		},
	}
}
