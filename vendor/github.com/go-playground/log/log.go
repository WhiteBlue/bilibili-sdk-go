package log

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// FilenameDisplay is the type of file display output
type FilenameDisplay uint8

// FilenameDisplay options
const (
	Lshortfile FilenameDisplay = iota
	Llongfile
)

// HandlerChannels is an array of handler channels
type HandlerChannels []chan<- *Entry

// LevelHandlerChannels is a group of Handler channels mapped by Level
type LevelHandlerChannels map[Level]HandlerChannels

// DurationFormatFunc is the function called for parsing Trace Duration
type DurationFormatFunc func(time.Duration) string

type logger struct {
	fieldPool        *sync.Pool
	entryPool        *sync.Pool
	tracePool        *sync.Pool
	channels         LevelHandlerChannels
	durationFunc     DurationFormatFunc
	timeFormat       string
	appID            string
	callerInfoLevels [9]bool
}

const (
	// DefaultTimeFormat is the default time format when parsing Time values.
	// it is exposed to allow handlers to use and not have to redefine
	DefaultTimeFormat = "2006-01-02T15:04:05.000000000Z07:00"
)

// Logger is the default instance of the log package
var (
	once            sync.Once
	Logger          *logger
	exitFunc        = os.Exit
	skipLevel       = 2
	stackTraceLimit = 7000
)

func init() {
	once.Do(func() {
		Logger = &logger{
			fieldPool: &sync.Pool{New: func() interface{} {
				return Field{}
			}},
			tracePool: &sync.Pool{New: func() interface{} {
				return new(TraceEntry)
			}},
			channels:     make(LevelHandlerChannels),
			durationFunc: func(d time.Duration) string { return d.String() },
			timeFormat:   DefaultTimeFormat,
			callerInfoLevels: [9]bool{
				true,
				false,
				false,
				false,
				true,
				true,
				true,
				true,
				true,
			},
		}

		Logger.entryPool = &sync.Pool{New: func() interface{} {
			return &Entry{
				wg:            new(sync.WaitGroup),
				ApplicationID: Logger.getApplicationID(),
			}
		}}
	})
}

// LeveledLogger interface for logging by level
type LeveledLogger interface {
	Debug(v ...interface{})
	Trace(v ...interface{}) Traceable
	Info(v ...interface{})
	Notice(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Panic(v ...interface{})
	Alert(v ...interface{})
	Fatal(v ...interface{})
	Debugf(msg string, v ...interface{})
	Tracef(msg string, v ...interface{}) Traceable
	Infof(msg string, v ...interface{})
	Noticef(msg string, v ...interface{})
	Warnf(msg string, v ...interface{})
	Errorf(msg string, v ...interface{})
	Panicf(msg string, v ...interface{})
	Alertf(msg string, v ...interface{})
	Fatalf(msg string, v ...interface{})
	WithFields(...Field) LeveledLogger
	StackTrace() LeveledLogger
}

var _ LeveledLogger = Logger

// Debug level formatted message.
func (l *logger) Debug(v ...interface{}) {
	e := newEntry(DebugLevel, fmt.Sprint(v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Trace starts a trace & returns Traceable object to End + log.
// Example defer log.Trace(...).End()
func (l *logger) Trace(v ...interface{}) Traceable {

	t := l.tracePool.Get().(*TraceEntry)
	t.entry = newEntry(TraceLevel, fmt.Sprint(v...), make([]Field, 0), skipLevel)
	t.start = time.Now().UTC()

	return t
}

// Info level formatted message.
func (l *logger) Info(v ...interface{}) {
	e := newEntry(InfoLevel, fmt.Sprint(v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Notice level formatted message.
func (l *logger) Notice(v ...interface{}) {
	e := newEntry(NoticeLevel, fmt.Sprint(v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Warn level formatted message.
func (l *logger) Warn(v ...interface{}) {
	e := newEntry(WarnLevel, fmt.Sprint(v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Error level formatted message.
func (l *logger) Error(v ...interface{}) {
	e := newEntry(ErrorLevel, fmt.Sprint(v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Panic logs an Panic level formatted message and then panics
func (l *logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	e := newEntry(PanicLevel, s, nil, skipLevel)
	l.HandleEntry(e)
	panic(s)
}

// Alert logs an Alert level formatted message and then panics
func (l *logger) Alert(v ...interface{}) {
	s := fmt.Sprint(v...)
	e := newEntry(AlertLevel, s, nil, skipLevel)
	l.HandleEntry(e)
}

// Fatal level formatted message, followed by an exit.
func (l *logger) Fatal(v ...interface{}) {
	e := newEntry(FatalLevel, fmt.Sprint(v...), nil, skipLevel)
	l.HandleEntry(e)
	exitFunc(1)
}

// Debugf level formatted message.
func (l *logger) Debugf(msg string, v ...interface{}) {
	e := newEntry(DebugLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Tracef starts a trace & returns Traceable object to End + log
func (l *logger) Tracef(msg string, v ...interface{}) Traceable {

	t := l.tracePool.Get().(*TraceEntry)
	t.entry = newEntry(TraceLevel, fmt.Sprintf(msg, v...), make([]Field, 0), skipLevel)
	t.start = time.Now().UTC()

	return t
}

// Infof level formatted message.
func (l *logger) Infof(msg string, v ...interface{}) {
	e := newEntry(InfoLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Noticef level formatted message.
func (l *logger) Noticef(msg string, v ...interface{}) {
	e := newEntry(NoticeLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Warnf level formatted message.
func (l *logger) Warnf(msg string, v ...interface{}) {
	e := newEntry(WarnLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Errorf level formatted message.
func (l *logger) Errorf(msg string, v ...interface{}) {
	e := newEntry(ErrorLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	l.HandleEntry(e)
}

// Panicf logs an Panic level formatted message and then panics
func (l *logger) Panicf(msg string, v ...interface{}) {
	s := fmt.Sprintf(msg, v...)
	e := newEntry(PanicLevel, s, nil, skipLevel)
	l.HandleEntry(e)
	panic(s)
}

// Alertf logs an Alert level formatted message and then panics
func (l *logger) Alertf(msg string, v ...interface{}) {
	s := fmt.Sprintf(msg, v...)
	e := newEntry(AlertLevel, s, nil, skipLevel)
	l.HandleEntry(e)
}

// Fatalf level formatted message, followed by an exit.
func (l *logger) Fatalf(msg string, v ...interface{}) {
	e := newEntry(FatalLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	l.HandleEntry(e)
	exitFunc(1)
}

// F creates a new field key + value entry
func (l *logger) F(key string, value interface{}) Field {

	fld := Logger.fieldPool.Get().(Field)
	fld.Key = key
	fld.Value = value

	return fld
}

// WithFields returns a log Entry with fields set
func (l *logger) WithFields(fields ...Field) LeveledLogger {
	return newEntry(InfoLevel, "", fields, skipLevel)
}

// StackTrace creates a new log Entry with pre-populated field with stack trace.
func (l *logger) StackTrace() LeveledLogger {
	trace := make([]byte, 1<<16)
	n := runtime.Stack(trace, true)
	if n > stackTraceLimit {
		n = stackTraceLimit
	}
	return newEntry(DebugLevel, "", []Field{F("stack trace", string(trace[:n])+"\n")}, skipLevel)
}

func (l *logger) HandleEntry(e *Entry) {

	// gather info if WarnLevel, ErrorLevel, PanicLevel, AlertLevel, FatalLevel or
	// gathering info for all levels, but only if no line info already exists; we could
	// be doing central logging.
	if e.Line == 0 && l.callerInfoLevels[e.Level] {
		_, e.File, e.Line, _ = runtime.Caller(e.calldepth)
	}

	//																											  ---------
	//																								|----------> | console |
	// Addding this check for when you are doing centralized logging                                |             ---------
	// i.e. -----------------               -----------------                                 -------------       --------
	//     | app log handler | -- json --> | central log app | -- unmarshal json to Entry -> | log handler | --> | syslog |
	//      -----------------               -----------------                                 -------------       --------
	//      																						|             ---------
	//      																						|----------> | DataDog |
	//      																									  ---------
	if e.wg == nil {
		e.wg = new(sync.WaitGroup)
	}

	channels, ok := l.channels[e.Level]
	if ok {

		e.wg.Add(len(channels))

		for _, ch := range channels {
			ch <- e
		}

		e.wg.Wait()
	}

	// reclaim entry + fields
	for _, f := range e.Fields {
		l.fieldPool.Put(f)
	}

	l.entryPool.Put(e)
}

// RegisterHandler adds a new Log Handler and specifies what log levels
// the handler will be passed log entries for
func (l *logger) RegisterHandler(handler Handler, levels ...Level) {

	ch := handler.Run()

	for _, level := range levels {

		channels, ok := l.channels[level]
		if !ok {
			channels = make(HandlerChannels, 0)
		}

		l.channels[level] = append(channels, ch)
	}

}

// RegisterDurationFunc registers a custom duration function for Trace events
func (l *logger) RegisterDurationFunc(fn DurationFormatFunc) {
	l.durationFunc = fn
}

// SetTimeFormat sets the time format used for Trace events
func (l *logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

// SetApplicationID tells the logger to set a constant application key
// that will be set on all log Entry objects. log does not care what it is,
// the application name, app name + hostname.... that's up to you
// it is needed by many logging platforms for separating logs by application
// and even by application server in a distributed app.
func (l *logger) SetApplicationID(id string) {
	l.appID = id
}

// SetCallerInfoLevels tells the logger to gather and set file and line number
// information on Entry objects for the provided log levels.
// By defaut all but TraceLevel, InfoLevel and NoticeLevel are set to gather information.
func (l *logger) SetCallerInfoLevels(levels ...Level) {
	for i := 0; i < len(l.callerInfoLevels); i++ {
		l.callerInfoLevels[i] = false
	}

	for i := 0; i < len(levels); i++ {
		l.callerInfoLevels[int(levels[i])] = true
	}
}

// SetCallerSkipDiff adds the provided diff to the caller SkipLevel values.
// This is used when wrapping this library, you can set ths to increase the
// skip values passed to Caller that retrieves the file + line number info.
func (l *logger) SetCallerSkipDiff(diff uint8) {
	skipLevel += int(diff)
}

// HasHandlers returns if any handlers have been registered.
func (l *logger) HasHandlers() bool {
	return len(l.channels) != 0
}

func (l *logger) getApplicationID() string {
	return l.appID
}
