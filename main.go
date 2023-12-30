package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyData struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api/endpoint", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		var data MyData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON message", http.StatusBadRequest)
			return
		}

		if data.Message == "" {
			http.Error(w, "Invalid JSON message", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received message: %s\n", data.Message)

		response := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Данные успешно приняты",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
