package main

import "fmt"

func main() {
	// err := get()
	// traceableErr := getWithTrace()
	traceableTypedErr := getWithErrorType()

	// PrintError(err)
	// PrintError(traceableErr)
	PrintError(traceableTypedErr)
}

func PrintError(err error) {
	fmt.Printf("\nError Type -> %T", err)
	fmt.Printf("\nError Message -> %s", err)
	fmt.Printf("\nTrace -> %+v", err)
}
