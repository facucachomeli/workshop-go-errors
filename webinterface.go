package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func get() error {
	_, err := findDomainEntity()
	if err != nil {
		return errors.New(fmt.Sprintf(notFoundMessage+": %s", err))
	}

	return nil
}

func getWithTrace() error {
	_, err := findDomainEntityWithTrace()
	return errors.Wrap(err, notFoundMessage)
}

func getWithErrorType() error {
	_, err := findDomainEntityWithErrorType()
	return NotFound{TraceableError{errors.Wrap(err, notFoundMessage)}}
}

const notFoundMessage = "404 - Not Found"

var NotFoundError = fmt.Errorf(notFoundMessage)

type NotFound struct {
	TraceableError
}
