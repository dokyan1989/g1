package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteOption ...
type WriteOption func(w http.ResponseWriter)

// WithResponseHeader ...
func WithResponseHeader(key string, value string) WriteOption {
	return func(w http.ResponseWriter) {
		w.Header().Set(key, value)
	}
}

// Write ...
func Write(w http.ResponseWriter, body interface{}, code int, opts ...WriteOption) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	for i := 0; i < len(opts); i++ {
		opt := opts[i]
		opt(w)
	}
	w.WriteHeader(code)
	switch v := body.(type) {
	case string:
		fmt.Fprint(w, v)
	case []byte:
		fmt.Fprint(w, string(v))
	default:
		fmt.Fprint(w, v)
	}
}

// WriteSuccess ...
func WriteSuccess(w http.ResponseWriter, data interface{}, statusCode int) {
	withJSONType := WithResponseHeader("Content-Type", "application/json")
	body, err := json.Marshal(data)
	if err != nil {
		Write(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Write(w, body, statusCode, withJSONType)
}

// WriteError ...
func WriteError(w http.ResponseWriter, err error) {
	msg := "Unknown error"
	code := http.StatusInternalServerError
	withJSONType := WithResponseHeader("Content-Type", "application/json")

	if err != nil {
		switch e := err.(type) {
		case Error:
			msg = e.Error()
			code = e.Status()

		default:
			msg = e.Error()
		}
	}

	body, err2 := json.Marshal(ErrorResponse{Message: msg})
	if err2 != nil {
		Write(w, err2.Error(), code)
		return
	}
	Write(w, body, code, withJSONType)

}
