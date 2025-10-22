package main

import (
	"fmt"
	_ "gov2panel/internal/packed"
	"gov2panel/internal/service"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "gov2panel/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gov2panel/internal/cmd"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
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
	}()

	cmd.Main.Run(gctx.GetInitCtx())
}
