package main

import "fmt"

// Returns customized hello message
func Hello(name string) string {
	rval := "Hello, World!"
	if name != "" {
		rval = fmt.Sprintf("Hello, %s!", name)
	}
	return rval
}

