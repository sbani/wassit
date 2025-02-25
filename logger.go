package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type Logger struct {
	infoOutput io.Writer
	errOutput  io.Writer
}

func NewSimpleLogger(quiet bool) Logger {
	var errOut, infoOut io.Writer
	if quiet {
		errOut, infoOut = ioutil.Discard, ioutil.Discard
	} else {
		errOut, infoOut = os.Stderr, os.Stdout
	}

	return Logger{
		infoOutput: infoOut,
		errOutput:  errOut,
	}
}

func (l Logger) Errorf(format string, in ...interface{}) {
	l.Error(fmt.Sprintf(format, in...))
}

func (l Logger) Error(msg string) {
	l.print(l.errOutput, msg)
}

func (l Logger) Critical(msg string) {
	l.print(l.errOutput, msg)
	os.Exit(1)
}

func (l Logger) Infof(format string, in ...interface{}) {
	l.Info(fmt.Sprintf(format, in...))
}

func (l Logger) Info(msg string) {
	l.print(l.infoOutput, msg)
}

func (l Logger) print(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "[%s] %s\n", time.Now().Format(time.RFC3339), msg)
}
