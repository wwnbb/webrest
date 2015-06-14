package webrest

import (
	"fmt"
)

func Logger(e Env) {

	// Request logger funcition
	// Prints request information to stdout

	fmt.Printf(
		"[%s] -- %s \"%s %s\" %v %d \n",
		e.Request.RemoteAddr,
		e.Request.Method,
		e.Request.URL,
		e.Request.ContentLength,
	)
}
