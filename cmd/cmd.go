package cmd

import (
	serviced "github.com/kardianos/service"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	sv "sb-cli/cmd/service"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

var (
	c       Config
	isCfgOk bool
	cfgFile string
)

func init() {

	// 加载运行的初始化配置
	rootCmd.PersistentFlags().StringVar(&cfgFile, "f", "cfg.yaml", "config file path")

	if err := conf.Load(cfgFile, &c); err != nil {
		logx.Errorf("error: config file %s, %s", cfgFile, err.Error())
	} else {
		isCfgOk = true
	}
	if err := logx.SetUp(c.Log); err != nil {
		logx.Must(err)
	}
	rootCmd.AddCommand(SvcCmd, initCmd)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "启动服务进程",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if !isCfgOk {
			os.Exit(-1)
		}
		serviceGroup := service.NewServiceGroup()
		//TODO: 考虑用 wired 注入
		// TODO:这里跑服务
		serviceGroup.Start()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

		s := <-quit

		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logx.Info("接收到退出信号")
			serviceGroup.Stop()
			logx.Info("服务完全退出")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// testCmd represents the stu command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化",

	Run: func(cmd *cobra.Command, args []string) {
		// 这里考虑生成默认的配置
		// 回头还要添加更多的方法来解决配置的修改和生成
	},
}

// SvcCmd  represents the stu command
var SvcCmd = &cobra.Command{
	Use:   "service",
	Short: "系统服务设置",
	Args:  cobra.ExactArgs(0),
}

func init() {
	_path, err := os.Executable()
	if err != nil {
		logx.Must(err)
	}
	dir := filepath.Dir(_path)
	svcConfig := &serviced.Config{
		Name:             "sb-helper",
		DisplayName:      "sb-helper service",
		Description:      "sing-box 服务",
		WorkingDirectory: dir,
		UserName:         "root",
		Option: map[string]interface{}{
			"LimitNOFILE": 65535,
			"RestartSec":  20, // TODO 暂不支持
		},
	}

	prg := &program{}
	s, err := serviced.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	SvcCmd.AddCommand(sv.NewSvcInstallCmd(s),
		sv.NewSvcUnInstallCmd(s),
		sv.NewSvcStartCmd(s), sv.NewSvcReStartCmd(s), sv.NewSvcStopCmd(s))
}

type program struct{}

func (p *program) Start(s serviced.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	logx.Infof("服务启动")
	Execute()
}

func (p *program) Stop(s serviced.Service) error {
	// Stop should not block. Return with a few seconds.
	logx.Infof("服务停止")
	return nil
}
