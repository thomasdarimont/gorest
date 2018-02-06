package actuator

import (
	"net/http"
	"encoding/json"
)

type Health struct {
	Status string `json: status`
}

func HealthEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(&Health{Status: "OK"})
}
