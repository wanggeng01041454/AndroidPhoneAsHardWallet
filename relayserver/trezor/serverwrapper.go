package trezor

import (
	"os"

	"github.com/gerik/aphw/relayserver/trezor/core"
	"github.com/gerik/aphw/relayserver/trezor/memorywriter"
	"github.com/gerik/aphw/relayserver/trezor/server"
	"github.com/gerik/aphw/relayserver/ui"
)

type TrezorServer interface {
	// 开始服务，
	// 开启 go routine, 不阻塞主线程
	StartNoHup() error
}

type trezorServer struct {
	// ui 界面的接口
	ui ui.RelayServerUI

	// 日志记录器
	log memorywriter.LogWriter

	// 网络设备管理
	netDevManager core.NetDeviceManager

	// 对接 metamask 的server
	metamaskServer *server.Server
}

var _ TrezorServer = (*trezorServer)(nil)

func CreateTrezorServer(ui ui.RelayServerUI) (TrezorServer, error) {
	log := newLogWrapper(ui)

	netDevManager := core.NewNetDevCore(ui)

	// todo, 改为dumy输出
	stderrWriter := os.Stdout
	// todo version 和 githash 模拟 2.0.4 版本的值
	version := "0.0.0"
	githash := "00000"

	server, err := server.New(netDevManager,
		stderrWriter,
		log,
		log,
		version,
		githash,
	)

	if err != nil {
		return nil, err
	}

	return &trezorServer{
		ui:             ui,
		log:            log,
		netDevManager:  netDevManager,
		metamaskServer: server,
	}, nil
}

// 启动服务
func (t *trezorServer) StartNoHup() error {
	go func() {
		t.metamaskServer.Run()
	}()

	return nil
}
