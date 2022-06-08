package log

import (
	"io"
	"os"
)

type Writer interface {
	Write(p []byte, flush bool) (n int, err error)
}

type writer struct {
	io.Writer
}

func (w writer) Write(p []byte, flush bool) (n int, err error) {
	return w.Writer.Write(p)
}

func NewWriter(w io.Writer) Writer {
	return writer{w}
}

type consoleWriter struct {
}

// 返回控制台输出数据的长度
func (w consoleWriter) Write(p []byte, flush bool) (n int, err error) {
	n, err = os.Stdout.Write(p)
	if flush {
		os.Stdout.Sync()
	}
	return
}

var ConsoleWriter Writer = consoleWriter{}

type discardWriter struct {
}

func (w discardWriter) Write(p []byte, flush bool) (n int, err error) {
	return len(p), nil
}

var DiscardWriter Writer = discardWriter{}
