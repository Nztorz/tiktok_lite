package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON[T any](r *http.Request) (T, error) {
	var data T

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)
	if err != nil {
		return data, fmt.Errorf("parsing json %w", err)
	}

	return data, nil
}
