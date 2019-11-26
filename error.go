package main

import (
	"fmt"
	"io"
)

type TraceableError struct {
	cause error
}

func (err TraceableError) Error() string {
	return err.cause.Error()
}

func (err TraceableError) Cause() error {
	return err.cause
}

func (err TraceableError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", err.Cause())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, err.Error())
	case 'q':
		fmt.Fprintf(s, "%q", err.Error())
	}
}
