package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var DomainEntityNotFoundError = fmt.Errorf(domainEntityNotFoundMessage)

const domainEntityNotFoundMessage = "Domain Entity Not Found"

func findDomainEntity() (interface{}, error) {
	document, err := find()
	return document, errors.New(fmt.Sprintf(domainEntityNotFoundMessage+": %v", err))
}

func findDomainEntityWithTrace() (interface{}, error) {
	document, err := findWithTrace()
	return document, errors.Wrap(err, domainEntityNotFoundMessage)
}

func findDomainEntityWithErrorType() (interface{}, error) {
	document, err := findWithErrorType()
	return document, DomainEntityNotFound{TraceableError{errors.Wrap(err, domainEntityNotFoundMessage)}}
}

type DomainEntityNotFound struct {
	TraceableError
}
