package memorywriter

import (
	"io"

	trezord_memorywriter "github.com/trezor/trezord-go/memorywriter"
)

// 定义一个和 memorywriter 暴漏接口一致的接口
type LogWriter interface {
	io.Writer

	// 记录日志
	Log(s string)

	// exprot as gzip
	// todo, 根据业务逻辑确认，这个是否有必要
	Gzip(start string) ([]byte, error)

	// export as string
	// todo, 根据业务逻辑确认，这个是否有必要
	String(start string) (string, error)
}

var _ LogWriter = (*trezord_memorywriter.MemoryWriter)(nil)
