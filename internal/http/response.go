package http

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, payload any) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func ResponseError(w http.ResponseWriter, code int, msg string) error {
	return ResponseJSON(w, code, map[string]string{"error": msg})
}
