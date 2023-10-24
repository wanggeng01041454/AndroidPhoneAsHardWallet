package trezor

import (
	"fmt"
	"io"
	"os"

	"github.com/gerik/aphw/relayserver/trezor/memorywriter"
	"github.com/gerik/aphw/relayserver/ui"
)

type logwrapper struct {
	ui     ui.RelayServerUI
	stdout io.Writer
}

var _ memorywriter.LogWriter = (*logwrapper)(nil)

func newLogWrapper(ui ui.RelayServerUI) memorywriter.LogWriter {
	return &logwrapper{
		ui:     ui,
		stdout: os.Stdout,
	}
}

func (u *logwrapper) Log(s string) {
	// todo
	u.stdout.Write([]byte(s))
}

func (u *logwrapper) Write(p []byte) (n int, err error) {
	// todo
	u.stdout.Write(p)
	return len(p), nil
}

func (u *logwrapper) Gzip(start string) ([]byte, error) {
	// todo
	return nil, fmt.Errorf("not implement")
}

func (u *logwrapper) String(start string) (string, error) {
	// todo
	return "", fmt.Errorf("not implement")
}
