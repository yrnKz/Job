package log

import "fmt"

type Level uint8

// 日志等级判断接口
type LevelEnabler interface {
	Enabled(Level) bool
}

const (
	LevelNone Level = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
)

const (
	levelNoneString  = "None"
	levelFatalString = "Fatal"
	levelErrorString = "Error"
	levelWarnString  = "Warn"
	levelInfoString  = "Info"
	levelDebugString = "Debug"
	levelTraceString = "Trace"

	levelNoneLowercaseString  = "none"
	levelFatalLowercaseString = "fatal"
	levelErrorLowercaseString = "error"
	levelWarnLowercaseString  = "warn"
	levelInfoLowercaseString  = "info"
	levelDebugLowercaseString = "debug"
	levelTraceLowercaseString = "trace"

	levelNoneUppercaseString  = "NONE"
	levelFatalUppercaseString = "FATAL"
	levelErrorUppercaseString = "ERROR"
	levelWarnUppercaseString  = "WARN"
	levelInfoUppercaseString  = "INFO"
	levelDebugUppercaseString = "DEBUG"
	levelTraceUppercaseString = "TRACE"

	UnknownLevelFormat          = "Level(%d)"
	UnknownLevelLowercaseFormat = "level(%d)"
	UnknownLevelUppercaseFormat = "LEVEL(%d)"
)

var (
	levelNoneColorString  = Red.Add(levelNoneUppercaseString)
	levelFatalColorString = Red.Add(levelFatalUppercaseString)
	levelErrorColorString = Red.Add(levelErrorUppercaseString)
	levelWarnColorString  = Yellow.Add(levelWarnUppercaseString)
	levelInfoColorString  = Green.Add(levelInfoUppercaseString)
	levelDebugColorString = White.Add(levelDebugUppercaseString)
	levelTraceColorString = White.Add(levelTraceUppercaseString)
)

var Levels = [...]Level{LevelFatal, LevelError, LevelWarn, LevelInfo, LevelDebug, LevelTrace}

func (l Level) Valid() bool {
	return LevelNone <= l && l <= LevelTrace
}

func (l Level) Enabled(lvl Level) bool {
	return lvl <= l
}

func (l Level) String() string {
	switch l {
	case LevelNone:
		return levelNoneString
	case LevelFatal:
		return levelFatalString
	case LevelError:
		return levelErrorString
	case LevelWarn:
		return levelWarnString
	case LevelInfo:
		return levelInfoString
	case LevelDebug:
		return levelDebugString
	case LevelTrace:
		return levelTraceString
	default:
		return fmt.Sprintf(UnknownLevelFormat, uint8(l))
	}
}

func (l Level) ColorString() string {
	switch l {
	case LevelNone:
		return levelNoneColorString
	case LevelFatal:
		return levelFatalColorString
	case LevelError:
		return levelErrorColorString
	case LevelWarn:
		return levelWarnColorString
	case LevelInfo:
		return levelInfoColorString
	case LevelDebug:
		return levelDebugColorString
	case LevelTrace:
		return levelTraceColorString
	default:
		return Red.Add(fmt.Sprintf(UnknownLevelUppercaseFormat, uint8(l)))
	}
}

// 强转日志等级
func ParseLevel(s string) (l Level, err error) {
	switch s {
	case levelNoneString, levelNoneLowercaseString, levelNoneUppercaseString:
		l = LevelNone
	case levelFatalString, levelFatalLowercaseString, levelFatalUppercaseString:
		l = LevelFatal
	case levelErrorString, levelErrorLowercaseString, levelErrorUppercaseString:
		l = LevelError
	case levelWarnString, levelWarnLowercaseString, levelWarnUppercaseString:
		l = LevelWarn
	case levelInfoString, levelInfoLowercaseString, levelInfoUppercaseString:
		l = LevelInfo
	case levelDebugString, levelDebugLowercaseString, levelDebugUppercaseString:
		l = LevelDebug
	case levelTraceString, levelTraceLowercaseString, levelTraceUppercaseString:
		l = LevelTrace
	default:
		var val uint8
		_, err = fmt.Sscanf(s, UnknownLevelFormat, &val)
		if err == nil {
			l = Level(val)
		}
	}
	return
}

// 日志格式
func (l Level) Format(s fmt.State, c rune) {
	switch l {
	case LevelNone:
		fmt.Fprint(s, levelNoneString)
	case LevelFatal:
		fmt.Fprint(s, levelFatalString)
	case LevelError:
		fmt.Fprint(s, levelErrorString)
	case LevelWarn:
		fmt.Fprint(s, levelWarnString)
	case LevelInfo:
		fmt.Fprint(s, levelInfoString)
	case LevelDebug:
		fmt.Fprint(s, levelDebugString)
	case LevelTrace:
		fmt.Fprint(s, levelTraceString)
	default:
		fmt.Fprintf(s, UnknownLevelFormat, uint8(l))
	}
}

func (l *Level) UnmarshalText(b []byte) error {
	var err error
	var lv Level
	if lv, err = ParseLevel(string(b)); err != nil {
		return err
	}
	*l = lv
	return nil
}

func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}
