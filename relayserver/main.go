package main

import (
	"github.com/gerik/aphw/relayserver/ui"
)

func main() {
	ui := ui.NewMainWindow("Wallet Relay Server")

	ui.StartUI()
}
