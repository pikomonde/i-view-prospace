package log

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

// TODO: if this package is used on other projects (>3 projects), move this to another repository

// Fields type, used to pass to `WithFields`.
type Fields log.Fields

// Error is used to log error with error file, function, and line. It also returns the error back.
func Error(fields Fields, err error) error {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	frame, _ := runtime.CallersFrames(pc).Next()
	log.WithFields(log.Fields(fields)).Errorf("%s:%d %s: %v\n", frame.File, frame.Line, frame.Function, err)
	return err
}

// Print implement logrus's logrus.Print
func Print(args ...interface{}) {
	log.Print(args...)
}

// Println implement logrus's logrus.Println
func Println(args ...interface{}) {
	log.Println(args...)
}

// Printf implement logrus's logrus.Printf
func Printf(args ...interface{}) {
	log.Println(args...)
}
