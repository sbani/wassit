package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type logger struct {
	infoOutput io.Writer
	errOutput  io.Writer
}

func simpleLogger(quiet bool) logger {
	var errOut, infoOut io.Writer
	if quiet {
		errOut, infoOut = ioutil.Discard, ioutil.Discard
	} else {
		errOut, infoOut = os.Stderr, os.Stdout
	}

	return logger{
		infoOutput: infoOut,
		errOutput:  errOut,
	}
}

func (l logger) Errorf(format string, in ...interface{}) {
	l.Error(fmt.Sprintf(format, in...))
}

func (l logger) Error(msg string) {
	l.print(l.errOutput, msg)
}

func (l logger) Critical(msg string) {
	l.print(l.errOutput, msg)
	os.Exit(1)
}

func (l logger) Infof(format string, in ...interface{}) {
	l.Info(fmt.Sprintf(format, in...))
}

func (l logger) Info(msg string) {
	l.print(l.infoOutput, msg)
}

func (l logger) print(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "[%s] %s\n", time.Now().Format(time.RFC3339), msg)
}
