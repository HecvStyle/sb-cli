package service

import (
	"fmt"

	"github.com/kardianos/service"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/spf13/cobra"
)

// NewSvcStartCmd  represents the stu command
func NewSvcStartCmd(s service.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "启动服务",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("启动服务")
			err := service.Control(s, "start")
			if err != nil {
				logx.Must(err)
			}
		},
	}
}
