package main

import (
	"fmt"

	"github.com/gerik/aphw/relayserver/trezor"
	"github.com/gerik/aphw/relayserver/ui"
)

func main() {
	ui := ui.NewMainWindow("Wallet Relay Server")

	// 创建服务
	server, err := trezor.CreateTrezorServer(ui)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_ = server.StartNoHup()

	ui.StartUI()
}
