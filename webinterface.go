package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func get() error {
	_, err := findDomainEntity()
	if err != nil {
		return NotFoundError
	}

	return nil
}

func getWithTrace() error {
	_, err := findDomainEntityWithTrace()
	return errors.Wrap(err, notFoundMessage)
}

func getWithErrorType() error {
	_, err := findDomainEntityWithErrorType()
	return NewNotFound(err)
}

const notFoundMessage = "404 - Not Found"

var NotFoundError = fmt.Errorf(notFoundMessage)

type NotFound struct {
	cause error
}

func NewNotFound(e error) error {
	return NotFound{errors.Wrap(e, notFoundMessage)}
}

func (err NotFound) Error() string {
	return err.cause.Error()
}

func (err NotFound) Cause() error {
	return err.cause
}

func (err NotFound) Format(s fmt.State, verb rune) {
	Format(err, s, verb)
}
