package main

import (
	"fmt"
	"net/http"
	"time"
)

// Handler function for the /datetime endpoint
func datetimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"current_datetime": "%s"}`, currentTime)
}

func main() {
	http.HandleFunc("/datetime", datetimeHandler)

	fmt.Println("Server is running at http://localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
