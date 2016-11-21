package log

// ANSIEscSeq is a predefined ANSI escape sequence
type ANSIEscSeq string

// ANSI escape sequences
// NOTE: in an standard xterm terminal the light colors will appear BOLD instead of the light variant
const (
	Black        ANSIEscSeq = "\x1b[30m"
	DarkGray                = "\x1b[30;1m"
	Blue                    = "\x1b[34m"
	LightBlue               = "\x1b[34;1m"
	Green                   = "\x1b[32m"
	LightGreen              = "\x1b[32;1m"
	Cyan                    = "\x1b[36m"
	LightCyan               = "\x1b[36;1m"
	Red                     = "\x1b[31m"
	LightRed                = "\x1b[31;1m"
	Magenta                 = "\x1b[35m"
	LightMagenta            = "\x1b[35;1m"
	Brown                   = "\x1b[33m"
	Yellow                  = "\x1b[33;1m"
	LightGray               = "\x1b[37m"
	White                   = "\x1b[37;1m"
	Underscore              = "\x1b[4m"
	Blink                   = "\x1b[5m"
	Inverse                 = "\x1b[7m"
	Reset                   = "\x1b[0m"
)
