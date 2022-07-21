package main

import (
	_ "fgf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"fgf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
