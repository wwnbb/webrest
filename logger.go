psf_admin-20150406.sql.gz
sf_front-20150406.sql.gzackage webrest

import (
	"fmt"
)

func Logger(e Env) {

	// Request logger funcition
	// Prints request information to stdout

	fmt.Printf(
		"[%s] -- %s \"%s %s\" %d %d \n",
		e.Request.Host,
		e.Request.RemoteAddr,
		e.Request.Method,
		e.Request.URL,
		e.Request.ContentLength,
	)
}
