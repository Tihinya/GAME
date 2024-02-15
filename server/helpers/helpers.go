package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"bomberman-dom/models"
)

type Broadcaster interface {
	BroadcastAllClients(event models.Event)
}

func ReturnMessageJSON(w http.ResponseWriter, message string, httpCode int, status string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(models.Response{
		Status:  status,
		Message: message,
	})
}

func SerializeData(EventType string, data ...any) models.Event {
	if len(data) == 1 {
		jsonData, err := json.Marshal(data[0])
		if err != nil {
			log.Printf("failed to marshal eventType %v: %v\n", EventType, err)
		}

		var outgoingEvent models.Event
		outgoingEvent.Payload = jsonData
		outgoingEvent.Type = EventType

		return outgoingEvent
	}
	return models.Event{}
}
