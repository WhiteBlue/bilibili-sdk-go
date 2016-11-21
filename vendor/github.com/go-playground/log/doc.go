/*
Package log is a simple, highly configurable, structured logging that is a near drop in replacement for the std library log.


Usage

    package main

    import (
        "github.com/go-playground/log"
        "github.com/go-playground/log/handlers/console"
    )

    func main() {

        cLog := console.New()

        log.RegisterHandler(cLog, log.AllLevels...)

        // Trace
        defer log.Trace("trace").End()

        log.Debug("debug")
        log.Info("info")
        log.Notice("notice")
        log.Warn("warn")
        log.Error("error")
        // log.Panic("panic") // this will panic
        log.Alert("alert")
        // log.Fatal("fatal") // this will call os.Exit(1)

        // logging with fields can be used with any of the above
        log.WithFields(log.F("key", "value")).Info("test info")
    }

Adding your own Handler

    package main

    import (
        "bytes"
        "fmt"

        "github.com/go-playground/log"
    )

    // CustomHandler is your custom handler
    type CustomHandler struct {
        // whatever properties you need
        buffer uint // channel buffer
    }

    // Run starts the logger consuming on the returned channed
    func (c *CustomHandler) Run() chan<- *log.Entry {

        // in a big high traffic app, set a higher buffer
        ch := make(chan *log.Entry, c.buffer)

        // can run as many consumers on the channel as you want,
        // depending on the buffer size or your needs
        go func(entries <-chan *log.Entry) {

            // below prints to os.Stderr but could marshal to JSON
            // and send to central logging server
            var e *log.Entry
            b := new(bytes.Buffer)

            for e = range entries {

                b.Reset()
                b.WriteString(e.Message)

                for _, f := range e.Fields {
                    fmt.Fprintf(b, " %s=%v", f.Key, f.Value)
                }

                fmt.Println(b.String())
                e.WG.Done() // done writing the entry
            }

        }(ch)

        return ch
    }

    func main() {

        cLog := &CustomHandler{
            buffer: 0,
        }

        log.RegisterHandler(cLog, log.AllLevels...)

        // Trace
        defer log.Trace("trace").End()

        log.Debug("debug")
        log.Info("info")
        log.Notice("notice")
        log.Warn("warn")
        log.Error("error")
        // log.Panic("panic") // this will panic
        log.Alert("alert")
        // log.Fatal("fatal") // this will call os.Exit(1)

        // logging with fields can be used with any of the above
        log.WithFields(log.F("key", "value")).Info("test info")
    }

Log Level Definitions

    DebugLevel - Info useful to developers for debugging the application, not useful during
                 operations.

    TraceLevel - Info useful to developers for debugging the application and reporting on
                 possible bottlenecks.

    InfoLevel - Normal operational messages - may be harvested for reporting, measuring
                throughput, etc. - no action required.

    NoticeLevel - Normal but significant condition. Events that are unusual but not error
                  conditions - might be summarized in an email to developers or admins to
                  spot potential problems - no immediate action required.

    WarnLevel - Warning messages, not an error, but indication that an error will occur if
                action is not taken, e.g. file system 85% full - each item must be resolved
                within a given time.

    ErrorLevel - Non-urgent failures, these should be relayed to developers or admins; each
                 item must be resolved within a given time.

    PanicLevel - A "panic" condition usually affecting multiple apps/servers/sites. At this
                 level it would usually notify all tech staff on call.

    AlertLevel - Action must be taken immediately. Should be corrected immediately,
                 therefore notify staff who can fix the problem. An example would be the
                 loss of a primary ISP connection.

    FatalLevel - Should be corrected immediately, but indicates failure in a primary
                 system, an example is a loss of a backup ISP connection.
                 ( same as SYSLOG CRITICAL )
*/
package log
