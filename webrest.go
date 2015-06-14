package webrest

import (
	"fmt"
	"net/http"
	"time"
)

func apphandler(r *http.Request) {
	e := Env{r, time.Now()}
	ViewSet(e)
	Logger(e)
}

func Run() {
	var localAddress string = ":8080"

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("WebRest » listening on %s...\n", localAddress)
	}()

	http.HandleFunc("/", apphandler)
	http.ListenAndServe(localAddress, nil)
}
