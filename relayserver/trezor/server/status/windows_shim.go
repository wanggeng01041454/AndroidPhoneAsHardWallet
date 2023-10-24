//go:build !windows
// +build !windows

// todo 删除这个文件，我们不需要这个文件

package status

import (
	"github.com/gerik/aphw/relayserver/trezor/memorywriter"
)

// Devcon is a tool for listing devices and drivers on windows
// These are empty functions that get called on *nix systems

func devconInfo(d memorywriter.LogWriter) (string, error) {
	return "", nil
}

func devconAllStatusInfo() (string, error) {
	return "", nil
}

func runMsinfo() (string, error) {
	return "", nil
}

func isWindows() bool {
	return false
}

func oldLog() (string, error) {
	return "", nil
}

func libwdiReinstallLog() (string, error) {
	return "", nil
}

func setupAPIDevLog() (string, error) {
	return "", nil
}
