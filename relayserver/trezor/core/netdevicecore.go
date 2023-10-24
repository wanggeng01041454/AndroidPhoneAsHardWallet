package core

import (
	"context"

	"github.com/gerik/aphw/relayserver/ui"
	trezordcore "github.com/trezor/trezord-go/core"
)

type NetDeviceManager interface {
	CoreInf

	// todo, 网络硬件钱包的管理方法
}

type netDeviceCore struct {
	ui ui.RelayServerUI
}

var _ NetDeviceManager = (*netDeviceCore)(nil)

func NewNetDevCore(ui ui.RelayServerUI) NetDeviceManager {
	return &netDeviceCore{
		ui: ui,
	}
}

func (n *netDeviceCore) Enumerate() ([]trezordcore.EnumerateEntry, error) {
	// todo
	return nil, nil
}

func (n *netDeviceCore) Listen(entries []trezordcore.EnumerateEntry, ctx context.Context) ([]trezordcore.EnumerateEntry, error) {
	// todo
	return nil, nil
}

func (n *netDeviceCore) Acquire(path, prev string, debug bool) (string, error) {
	// todo
	return "", nil
}

func (n *netDeviceCore) Release(session string, debug bool) error {
	// todo
	return nil
}

func (n *netDeviceCore) Call(body []byte, ssid string, mode trezordcore.CallMode, debug bool, ctx context.Context) ([]byte, error) {
	// todo
	return nil, nil
}
