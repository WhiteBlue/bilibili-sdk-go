package log

// AllLevels is an array of all log levels, for easier registering of all levels to a handler
var AllLevels = []Level{
	DebugLevel,
	TraceLevel,
	InfoLevel,
	NoticeLevel,
	WarnLevel,
	ErrorLevel,
	PanicLevel,
	AlertLevel,
	FatalLevel,
}

// Level of the log
type Level uint8

// Log levels.
const (
	DebugLevel Level = iota
	TraceLevel
	InfoLevel
	NoticeLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	AlertLevel
	FatalLevel // same as syslog CRITICAL
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case TraceLevel:
		return "TRACE"
	case InfoLevel:
		return "INFO"
	case NoticeLevel:
		return "NOTICE"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	case AlertLevel:
		return "ALERT"
	case FatalLevel:
		return "FATAL"
	default:
		return "Unknow Level"
	}
}

// TODO: Add a bytes method along with string
