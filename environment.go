package webrest

import (
	"net/http"
	"time"
)

type Env struct {
	Request *http.Request
	Time    time.Time
}
