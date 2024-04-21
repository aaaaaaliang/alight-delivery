package main

import (
	_ "alight-delivery/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"alight-delivery/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
