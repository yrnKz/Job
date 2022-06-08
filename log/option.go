package log

import "io"

type Option interface {
	apply(*logger)
}

type optionFunc func(*logger)

func (f optionFunc) apply(l *logger) {
	f(l)
}

func NowFuncOption(nowFunc NowFunc) Option {
	return optionFunc(func(l *logger) {
		l.nowFunc = nowFunc
	})
}

func AddCallerSkipOption(skip int) Option {
	return optionFunc(func(l *logger) {
		l.callerSkip += skip
	})
}

func AddCallerOption(addCaller bool) Option {
	return optionFunc(func(l *logger) {
		l.addCaller = addCaller
	})
}

func ErrorWriterOption(writer io.Writer) Option {
	return optionFunc(func(l *logger) {
		l.errorWriter = writer
	})
}
