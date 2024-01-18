package helpers

import (
	"encoding/json"
	"net/http"

	"bomberman-dom/models"
)

func ReturnMessageJSON(w http.ResponseWriter, message string, httpCode int, status string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(models.Response{
		Status:  status,
		Message: message,
	})
}
