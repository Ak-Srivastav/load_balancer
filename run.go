package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message    string `json:"message"`
	InstanceID string `json:"instance_id"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	instanceID := os.Getenv("INSTANCE_ID")
	if instanceID == "" {
		currentInstanceID := time.Now().String()
		err := os.Setenv("INSTANCE_ID", currentInstanceID)
		if err != nil {
			http.Error(w, "Failed to set instance ID", http.StatusInternalServerError)
			return
		}
		instanceID = currentInstanceID
	}
	log.Printf("Request received by instance: %s", instanceID)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Instance-ID", instanceID)
	response := Response{
		Message:    "Hello from instance",
		InstanceID: instanceID,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	instanceID := os.Getenv("INSTANCE_ID")
	if instanceID == "" {
		instanceID = time.Now().String()
		_ = os.Setenv("INSTANCE_ID", instanceID)
	}
	http.HandleFunc("/", handler)
	log.Printf("Server started at port: %d %s %s", 4001, " with instance ID: ", instanceID)
	log.Fatal(http.ListenAndServe(":4001", nil))
}
