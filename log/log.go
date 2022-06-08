package log

import "time"

var l = NewLogger(LevelDebug, ConsoleWriter, NowFuncOption(time.Now), AddCallerOption(true), AddCallerSkipOption(1))

func Set(v Logger) Logger {
	var oldV = l
	l = v
	return oldV
}

func Enabled(level Level) bool {
	return l.Enabled(level)
}

func Log(level Level, format string, args ...interface{}) {
	l.Log(level, format, args...)
}

// 致命日志
func Fatal(format string, args ...interface{}) {
	l.Log(LevelFatal, format, args...)
}

// 错误日志
func Error(format string, args ...interface{}) {
	l.Log(LevelError, format, args...)
}

// 警告日志
func Warn(format string, args ...interface{}) {
	l.Log(LevelWarn, format, args...)
}

// 信息日志
func Info(format string, args ...interface{}) {
	l.Log(LevelInfo, format, args...)
}

// 调试日志
func Debug(format string, args ...interface{}) {
	l.Log(LevelDebug, format, args...)
}

// 跟踪日志
func Trace(msg string, args ...interface{}) {
	l.Log(LevelTrace, msg, args...)
}
