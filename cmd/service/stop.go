package service

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/kardianos/service"

	"github.com/spf13/cobra"
)

func NewSvcStopCmd(s service.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "关闭服务",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("关闭服务")
			err := service.Control(s, "stop")
			if err != nil {
				logx.Must(err)
			}
		},
	}
}
