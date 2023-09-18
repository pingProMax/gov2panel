package main

import (
	_ "gov2panel/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "gov2panel/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gov2panel/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
