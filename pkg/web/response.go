package web

import (
	"encoding/json"
	"net/http"

	"github.com/marcelofabianov/fault"
)

func Success(w http.ResponseWriter, r *http.Request, status int, data any) {
	writeJSON(w, status, data)
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	response := fault.ToResponse(err)
	writeJSON(w, response.StatusCode, response)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func ErrDecodeRequestBody(err error, w http.ResponseWriter, r *http.Request) {
	Error(w, r, fault.NewInternalError(err, map[string]any{"reason": "failed to decode request body"}))
}

func ErrCreateAuditUser(err error, w http.ResponseWriter, r *http.Request) {
	Error(w, r, fault.NewInternalError(err, map[string]any{"reason": "failed to create audit user"}))
}
