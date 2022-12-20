package main

import (
	"strings"
)

// flagSlice is a type that implements the flag.Value interface.
type flagSlice []string

// String returns the string representation of the slice of strings.
func (f *flagSlice) String() string {
	return strings.Join(*f, ",")
}

// Set appends the string value to the slice of strings.
func (f *flagSlice) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func (f *flagSlice) IsEmpty() bool {
	return len(*f) == 0
}
