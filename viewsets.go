package webrest

import (
	"encoding/json"
	"net/http"
	"time"
)

type ViewSetError struct {
	Time    time.Time
	Message string
}

func Serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func ViewSet(e Env) error {
	str := ""
	_, err := Serialize(str)
	if err != nil {
		http.Error(e.Response, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}
