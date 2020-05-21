package log

import (
	"io"
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

// Setup sets up the logger.
func Setup(out io.Writer) {
	ldebug = log.New(out, "[DEBUG] ", log.Lshortfile)
	linfo = log.New(out, "[INFO] ", log.LstdFlags)
	lwarn = log.New(out, "[WARNING] ", log.LstdFlags)
	lfatal = log.New(out, "[ERROR] ", log.Lshortfile|log.LstdFlags)
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
