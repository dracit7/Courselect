package log

import (
	"bytes"
	"log"
	"os"
)

// Loggers with different types.
var (
	ldebug *log.Logger
	linfo  *log.Logger
	lwarn  *log.Logger
	lfatal *log.Logger
)

// Buffers for each logger.
var (
	bufdebug bytes.Buffer
	bufinfo  bytes.Buffer
	bufwarn  bytes.Buffer
	buffatal bytes.Buffer
)

// Setup sets up the logger.
func Setup() {
	ldebug = log.New(&bufdebug, "[DEBUG] ", log.Lshortfile)
	linfo = log.New(&bufinfo, "[INFO] ", log.Ldate)
	lwarn = log.New(&bufwarn, "[WARNING] ", log.Ldate)
	lfatal = log.New(&buffatal, "[ERROR] ", log.Lshortfile)
}

// Debug prints debug message with file/line information.
func Debug(s string) {
	ldebug.Output(2, s)
}

// Info prints info message.
func Info(s string) {
	linfo.Output(2, s)
}

// Warning prints warning message.
func Warning(s string) {
	lwarn.Output(2, s)
}

// Fatal prints error message and exit.
func Fatal(s string) {
	lfatal.Output(2, s)
	os.Exit(1)
}
