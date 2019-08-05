package main

import (
	"fmt"
	"io"
)

func Format(err Error, s fmt.State, verb rune) {
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

type Error interface {
	Error() string
	Cause() error
}
