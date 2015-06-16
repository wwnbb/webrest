package webrest

import (
	"net/http"
)

type Env struct {
	Request  *http.Request
	Response http.ResponseWriter
}
