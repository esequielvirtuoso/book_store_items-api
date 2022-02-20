package httputils

import (
	"encoding/json"
	"net/http"

	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err restErrors.RestErr) {
	RespondJson(w, err.Status(), err)
}
