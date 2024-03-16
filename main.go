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
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("程序关闭")
		service.User().MSaveAllRam()
		os.Exit(0)
	}()

	cmd.Main.Run(gctx.GetInitCtx())
}
