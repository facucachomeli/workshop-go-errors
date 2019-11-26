package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func find() (interface{}, error) {
	return nil, errors.New(documentNotFoundMessage)
}

func findWithTrace() (interface{}, error) {
	return nil, errors.New(documentNotFoundMessage)
}

func findWithErrorType() (interface{}, error) {
	return nil, DocumentNotFound{TraceableError{errors.New(documentNotFoundMessage)}}
}

const documentNotFoundMessage = "Document Not Found"

var DocumentNotFoundError = fmt.Errorf(documentNotFoundMessage)

type DocumentNotFound struct {
	TraceableError
}
