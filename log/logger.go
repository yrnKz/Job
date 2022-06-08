package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"github.com/pkg/errors"
)

// 日志接口
type Logger interface {
	LevelEnabler

	Log(level Level, format string, args ...interface{})

	Fatal(format string, args ...interface{})

	Error(format string, args ...interface{})

	Warn(format string, args ...interface{})

	Info(format string, args ...interface{})

	Debug(format string, args ...interface{})

	Trace(format string, args ...interface{})
}

var DiscardLogger Logger = discardLogger{}

type discardLogger struct {
}

func (d discardLogger) Enabled(level Level) bool {
	return false
}

func (d discardLogger) Log(level Level, format string, args ...interface{}) {

}

func (d discardLogger) Fatal(format string, args ...interface{}) {

}

func (d discardLogger) Error(format string, args ...interface{}) {

}

func (d discardLogger) Warn(format string, args ...interface{}) {

}

func (d discardLogger) Info(format string, args ...interface{}) {

}

func (d discardLogger) Debug(format string, args ...interface{}) {

}

func (d discardLogger) Trace(format string, args ...interface{}) {

}

type NowFunc func() time.Time

type logger struct {
	level       LevelEnabler
	writer      Writer
	nowFunc     NowFunc
	callerSkip  int
	addCaller   bool
	errorWriter io.Writer
}

func NewLogger(level LevelEnabler, writer Writer, options ...Option) Logger {
	if writer == nil {
		writer = DiscardWriter
	}
	var l = &logger{
		level:       level,
		writer:      writer,
		nowFunc:     nil,
		callerSkip:  0,
		addCaller:   false,
		errorWriter: os.Stderr,
	}
	var optionLen = len(options)
	for i := 0; i < optionLen; i++ {
		options[i].apply(l)
	}
	return l
}

func (l *logger) Enabled(level Level) bool {
	return l.level.Enabled(level)
}

func (l *logger) Log(level Level, format string, args ...interface{}) {
	l.log(level, format, args...)
}

// 致命日志
func (l *logger) Fatal(format string, args ...interface{}) {
	l.log(LevelFatal, format, args...)
}

// 错误日志
func (l *logger) Error(format string, args ...interface{}) {
	l.log(LevelError, format, args...)
}

// 警告日志
func (l *logger) Warn(format string, args ...interface{}) {
	l.log(LevelWarn, format, args...)
}

// 信息日志
func (l *logger) Info(format string, args ...interface{}) {
	l.log(LevelInfo, format, args...)
}

// 调试日志
func (l *logger) Debug(format string, args ...interface{}) {
	l.log(LevelDebug, format, args...)
}

// 跟踪日志
func (l *logger) Trace(msg string, args ...interface{}) {
	l.log(LevelTrace, msg, args...)
}

func (l *logger) logError(err error) {
	fmt.Fprintf(l.errorWriter, "error: %+v\n", err.Error())
}

func (l *logger) log(level Level, format string, args ...interface{}) {
	if !l.level.Enabled(level) {
		return
	}
	var err error
	var buf = &bytes.Buffer{}
	const callerSkipOffset = 2
	if l.nowFunc != nil {
		var t = l.nowFunc()
		if _, err = buf.WriteString(t.Format("2006-01-02T15:04:05.000000MST")); err != nil {
			l.logError(errors.WithStack(err))
			return
		}
		if err = buf.WriteByte(' '); err != nil {
			l.logError(errors.WithStack(err))
			return
		}
	}
	if _, err = fmt.Fprintf(buf, "[%s] ", level.ColorString()); err != nil {
		l.logError(errors.WithStack(err))
		return
	}
	if l.addCaller {
		if _, file, line, ok := runtime.Caller(l.callerSkip + callerSkipOffset); ok {
			fmt.Fprintf(buf, "%s:%d ", file, line)
		}
	}
	if _, err = fmt.Fprintf(buf, format, args...); err != nil {
		l.logError(errors.WithStack(err))
		return
	}
	if err = buf.WriteByte('\n'); err != nil {
		l.logError(errors.WithStack(err))
		return
	}
	if _, err = l.writer.Write(buf.Bytes(), LevelError.Enabled(level)); err != nil {
		l.logError(errors.WithStack(err))
		return
	}
}
