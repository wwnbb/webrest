package viewsets

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
	resp, err := Serialize(str)
	if err != nil {
		http.Error(e.Request.Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	e.Writer.Write(resp)
	return nil
}
