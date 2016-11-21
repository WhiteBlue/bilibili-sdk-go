/*
Package console allows for log messages to be sent to a any writer, default os.Stderr.

Example

simple console

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
*/
package console
