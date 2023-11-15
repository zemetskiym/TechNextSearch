package main

import (
    "strings"
    "net/http"
    "encoding/json"
)

type Response struct {
    Message string `json:"message"`
    Status int `json:"status"`
}

func main() {
    // Define port
    port := ":8080"
    
    // Handle calls for various routes
    http.HandleFunc("/api/", handleSearch)

    // Start an HTTP server on the given port
    http.ListenAndServe(port, nil)
}

func handleSearch (w http.ResponseWriter, r *http.Request) {
    // Check for query parameters
    queryParams := r.URL.Query()
    if len(queryParams) == 0 {
        http.Error(w, "Missing query parameters", http.StatusBadRequest)
        return
    }

    // Get a slice of values for the "query" key
    queryKeys := make([]string, 0, len(queryParams))
    for key, _ := range queryParams {
        queryKeys = append(queryKeys, key)
    }

    // Construct the response struct
    response := Response{
        Message: strings.Join(queryKeys, " "),
        Status: http.StatusOK,
    }

    // Convert the struct the JSON
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }
    
    // Set content type and write the JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}
