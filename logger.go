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

func newLogger(quiet bool) logger {
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
	log.Error(fmt.Sprintf(format, in))
}

func (l logger) Error(msg string) {
	log.print(l.errOutput, msg)
}

func (l logger) Critical(msg string) {
	log.print(l.errOutput, msg)
	os.Exit(1)
}

func (l logger) Infof(format string, in ...interface{}) {
	log.Info(fmt.Sprintf(format, in))
}

func (l logger) Info(msg string) {
	log.print(l.infoOutput, msg)
}

func (l logger) print(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "[%s] %s\n", time.Now().Format(time.RFC3339), msg)
}
