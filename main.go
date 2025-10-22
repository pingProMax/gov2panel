package main

import (
	"fmt"
	"gov2panel/internal/cmd"
	_ "gov2panel/internal/packed"
	"gov2panel/internal/service"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	_ "gov2panel/internal/logic"
)

func main() {

	go cmd.Main.Run(gctx.GetInitCtx())

	// 监听退出信号
	c := make(chan os.Signal, 1)
	// SIGHUP: terminal closed
	// SIGINT: Ctrl+C
	// SIGTERM: program exit
	// SIGQUIT: Ctrl+/
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	waitElegantExit(c)
}

// 优雅退出（退出信号）
func waitElegantExit(c chan os.Signal) {
	for i := range c {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("程序关闭")

			// 保存用户流量数据到数据库
			err := service.User().MSaveAllRam()
			if err != nil {
				fmt.Println("保存用户流量数据失败：", err.Error())
			} else {
				fmt.Println("保存用户流量数据成功")
			}

			//保存用户 7天流量使用数据 到文件
			err = service.User().SaveUserDay7FlowToFile(gctx.GetInitCtx(), "./resource/user_day7_flow.json")
			if err != nil {
				fmt.Println("保存用户流量数据失败：", err.Error())
			} else {
				fmt.Println("保存用户7天流量使用数据成功")

			}
			os.Exit(0)
		}
	}
}
