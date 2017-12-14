package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	Success interface{} `json:"success"`
}

func (s Success) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res, err := json.Marshal(s.Success)
	if err != nil {
		ServerError(w)
	}
	w.Write(res)
}

func ServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Write([]byte(`{"error": "unexpected error"}`))
}

func ClientError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Write([]byte(`{"error": "invalid request"}`))
}
