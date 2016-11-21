package log

import (
	"fmt"
	"runtime"
	"time"
)

// F creates a new field key + value entry
func F(key string, value interface{}) Field {
	return Logger.F(key, value)
}

// Debug level formatted message.
func Debug(v ...interface{}) {
	e := newEntry(DebugLevel, fmt.Sprint(v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Trace starts a trace & returns Traceable object to End + log.
// Example defer log.Trace(...).End()
func Trace(v ...interface{}) Traceable {
	t := Logger.tracePool.Get().(*TraceEntry)
	t.entry = newEntry(TraceLevel, fmt.Sprint(v...), make([]Field, 0), skipLevel)
	t.start = time.Now().UTC()

	return t
}

// Info level formatted message.
func Info(v ...interface{}) {
	e := newEntry(InfoLevel, fmt.Sprint(v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Notice level formatted message.
func Notice(v ...interface{}) {
	e := newEntry(NoticeLevel, fmt.Sprint(v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Warn level formatted message.
func Warn(v ...interface{}) {
	e := newEntry(WarnLevel, fmt.Sprint(v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Error level formatted message.
func Error(v ...interface{}) {
	e := newEntry(ErrorLevel, fmt.Sprint(v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Panic logs an Panic level formatted message and then panics
// it is here to let this log package be a drop in replacement
// for the standard logger
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	e := newEntry(PanicLevel, s, nil, skipLevel)
	Logger.HandleEntry(e)
	panic(s)
}

// Alert level formatted message.
func Alert(v ...interface{}) {
	s := fmt.Sprint(v...)
	e := newEntry(AlertLevel, s, nil, skipLevel)
	Logger.HandleEntry(e)
}

// Fatal level formatted message, followed by an exit.
func Fatal(v ...interface{}) {
	e := newEntry(FatalLevel, fmt.Sprint(v...), nil, skipLevel)
	Logger.HandleEntry(e)
	exitFunc(1)
}

// Fatalln level formatted message, followed by an exit.
func Fatalln(v ...interface{}) {
	e := newEntry(FatalLevel, fmt.Sprint(v...), nil, skipLevel)
	Logger.HandleEntry(e)
	exitFunc(1)
}

// Debugf level formatted message.
func Debugf(msg string, v ...interface{}) {
	e := newEntry(DebugLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Tracef starts a trace & returns Traceable object to End + log
func Tracef(msg string, v ...interface{}) Traceable {
	t := Logger.tracePool.Get().(*TraceEntry)
	t.entry = newEntry(TraceLevel, fmt.Sprintf(msg, v...), make([]Field, 0), skipLevel)
	t.start = time.Now().UTC()

	return t
}

// Infof level formatted message.
func Infof(msg string, v ...interface{}) {
	e := newEntry(InfoLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Noticef level formatted message.
func Noticef(msg string, v ...interface{}) {
	e := newEntry(NoticeLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Warnf level formatted message.
func Warnf(msg string, v ...interface{}) {
	e := newEntry(WarnLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Errorf level formatted message.
func Errorf(msg string, v ...interface{}) {
	e := newEntry(ErrorLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Panicf logs an Panic level formatted message and then panics
// it is here to let this log package be a near drop in replacement
// for the standard logger
func Panicf(msg string, v ...interface{}) {
	s := fmt.Sprintf(msg, v...)
	e := newEntry(PanicLevel, s, nil, skipLevel)
	Logger.HandleEntry(e)
	panic(s)
}

// Alertf level formatted message.
func Alertf(msg string, v ...interface{}) {
	s := fmt.Sprintf(msg, v...)
	e := newEntry(AlertLevel, s, nil, skipLevel)
	Logger.HandleEntry(e)
}

// Fatalf level formatted message, followed by an exit.
func Fatalf(msg string, v ...interface{}) {
	e := newEntry(FatalLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	Logger.HandleEntry(e)
	exitFunc(1)
}

// Panicln logs an Panic level formatted message and then panics
// it is here to let this log package be a near drop in replacement
// for the standard logger
func Panicln(v ...interface{}) {
	s := fmt.Sprint(v...)
	e := newEntry(PanicLevel, s, nil, skipLevel)
	Logger.HandleEntry(e)
	panic(s)
}

// Print logs an Info level formatted message
func Print(v ...interface{}) {
	e := newEntry(InfoLevel, fmt.Sprint(v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Println logs an Info level formatted message
func Println(v ...interface{}) {
	e := newEntry(InfoLevel, fmt.Sprint(v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// Printf logs an Info level formatted message
func Printf(msg string, v ...interface{}) {
	e := newEntry(InfoLevel, fmt.Sprintf(msg, v...), nil, skipLevel)
	Logger.HandleEntry(e)
}

// WithFields returns a log Entry with fields set
func WithFields(fields ...Field) LeveledLogger {
	return newEntry(InfoLevel, "", fields, skipLevel)
}

// StackTrace creates a new log Entry with pre-populated field with stack trace.
func StackTrace() LeveledLogger {
	trace := make([]byte, 1<<16)
	n := runtime.Stack(trace, true)
	if n > stackTraceLimit {
		n = stackTraceLimit
	}
	return newEntry(DebugLevel, "", []Field{F("stack trace", string(trace[:n])+"\n")}, skipLevel)
}

// HandleEntry send the logs entry out to all the registered handlers
func HandleEntry(e *Entry) {
	Logger.HandleEntry(e)
}

// RegisterHandler adds a new Log Handler and specifies what log levels
// the handler will be passed log entries for
func RegisterHandler(handler Handler, levels ...Level) {
	Logger.RegisterHandler(handler, levels...)
}

// RegisterDurationFunc registers a custom duration function for Trace events
func RegisterDurationFunc(fn DurationFormatFunc) {
	Logger.RegisterDurationFunc(fn)
}

// SetTimeFormat sets the time format used for Trace events
func SetTimeFormat(format string) {
	Logger.SetTimeFormat(format)
}

// SetCallerInfoLevels tells the logger to gather and set file and line number
// information on Entry objects for the provided log levels.
// By defaut all but TraceLevel, InfoLevel and NoticeLevel are set to gather information.
func SetCallerInfoLevels(levels ...Level) {
	Logger.SetCallerInfoLevels(levels...)
}

// SetCallerSkipDiff adds the provided diff to the caller SkipLevel values.
// This is used when wrapping this library, you can set ths to increase the
// skip values passed to Caller that retrieves the file + line number info.
func SetCallerSkipDiff(diff uint8) {
	Logger.SetCallerSkipDiff(diff)
}

// HasHandlers returns if any handlers have been registered.
func HasHandlers() bool {
	return Logger.HasHandlers()
}

// SetApplicationID tells the logger to set a constant application key
// that will be set on all log Entry objects. log does not care what it is,
// the application name, app name + hostname.... that's up to you
// it is needed by many logging platforms for separating logs by application
// and even by application server in a distributed app.
func SetApplicationID(id string) {
	Logger.SetApplicationID(id)
}
