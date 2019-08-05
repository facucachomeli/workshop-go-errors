package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func find() (interface{}, error) {
	return nil, DocumentNotFoundError
}

func findWithTrace() (interface{}, error) {
	return nil, errors.New(documentNotFoundMessage)
}

func findWithErrorType() (interface{}, error) {
	return nil, NewDocumentNotFound()
}

const documentNotFoundMessage = "Document Not Found"

var DocumentNotFoundError = fmt.Errorf(documentNotFoundMessage)

type DocumentNotFound struct {
	innerError error
}

func NewDocumentNotFound() error {
	return errors.New(documentNotFoundMessage)
}

func (err DocumentNotFound) Error() string {
	return documentNotFoundMessage
}

func (err DocumentNotFound) Cause() error {
	return err.innerError
}

func (err DocumentNotFound) Format(s fmt.State, verb rune) {
	Format(err, s, verb)
}
