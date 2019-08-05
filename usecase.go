package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var DomainEntityNotFoundError = fmt.Errorf(domainEntityNotFoundMessage)

const domainEntityNotFoundMessage = "Domain Entity Not Found"

func findDomainEntity() (interface{}, error) {
	document, err := find()
	return document, fmt.Errorf(domainEntityNotFoundMessage+": %v", err)
}

func findDomainEntityWithTrace() (interface{}, error) {
	document, err := findWithTrace()
	return document, errors.Wrap(err, domainEntityNotFoundMessage)
}

func findDomainEntityWithErrorType() (interface{}, error) {
	document, err := findWithErrorType()
	return document, NewDomainEntityNotFound(err)
}

type DomainEntityNotFound struct {
	cause      error
	innerError error
}

func NewDomainEntityNotFound(e error) error {
	return DomainEntityNotFound{errors.Wrap(e, domainEntityNotFoundMessage), e}
}

func (err DomainEntityNotFound) Error() string {
	return err.cause.Error()
}

func (err DomainEntityNotFound) Cause() error {
	return err.cause
}

func (err DomainEntityNotFound) Format(s fmt.State, verb rune) {
	Format(err, s, verb)
}
