package service

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

func NewSvcReStartCmd(s service.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "restart",
		Short: "启动服务",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("重启服务")
			err := service.Control(s, "restart")
			if err != nil {
				logx.Must(err)
			}
		},
	}
}
