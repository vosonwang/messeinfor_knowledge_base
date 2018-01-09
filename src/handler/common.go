package handler

import (
	"net/http"
	"encoding/json"
)

func JsonResponse(response interface{}, w http.ResponseWriter) {
	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}