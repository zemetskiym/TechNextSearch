package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	_ "github.com/mattn/go-sqlite3"
)

type Response struct {
    Message []Row `json:"message"`
    Status int `json:"status"`
}

type Row struct {
    Idx int
    Idx2 int
    PatentId int
    PatentText string
    Phase string
    Date string
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
    queryString := r.URL.Query().Get("query")
    
    // Convert the query string to be used to query database
    queryString = strings.ReplaceAll(queryString, " ", " + ")

    // Open a database connection
    db, err := sql.Open("sqlite3", "./database/patents.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Define the parameterized query to perform
    query := "SELECT * FROM virtual_patents_table WHERE patent_text MATCH ? ORDER BY rank;"

    // Perform match query on virtual table for full-text search
    rows, err := db.Query(query, queryString)
    if err != nil {
        log.Printf("Error querying database: %v", err)
        http.Error(w, "Error querying database", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Initialize slice to hold row data
    var results []Row

    // Iterate over the rows
    for rows.Next() {
        // Create variables to store column values
        var Idx, Idx2, PatentId int
        var PatentText, Phase, Date string

        // Scan the values from the row into variables
        err := rows.Scan(&Idx, &Idx2, &PatentId, &PatentText, &Phase, &Date)
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            http.Error(w, "Error scanning row", http.StatusInternalServerError)
            return
        }

        // Create a struct instance and populate it with the values
        result := Row{
            Idx: Idx,
            Idx2: Idx2,
            PatentId: PatentId,
            PatentText: PatentText,
            Phase: Phase,
            Date: Date,
        }

        // Append the struct to the results data
        results = append(results, result)
    }
    
    // Construct the response struct
    response := Response{
        Message: results,
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
