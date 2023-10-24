package core

import (
	"context"

	trezordcore "github.com/trezor/trezord-go/core"
)

type CoreInf interface {
	// 枚举设备
	Enumerate() ([]trezordcore.EnumerateEntry, error)

	// todo, 未知
	Listen(entries []trezordcore.EnumerateEntry, ctx context.Context) ([]trezordcore.EnumerateEntry, error)

	Acquire(path, prev string, debug bool) (string, error)

	Release(session string, debug bool) error

	Call(body []byte, ssid string, mode trezordcore.CallMode, debug bool, ctx context.Context) ([]byte, error)
}
