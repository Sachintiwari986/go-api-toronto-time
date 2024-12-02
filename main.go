package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Struct for API Response
type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

var db *sql.DB

// database connection
func connDB() {
	var err error
	// database connection strings
	dsn := "root:@tcp(127.0.0.1:3306)/timedb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error: Unable to open database connection: %v\n", err)
	}

	// Test database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error: Unable to connect to database: %v\n", err)
	}
	log.Println("Database connection successful!")
}

// Push data into the database
func DatatoDatabase(timestamp time.Time) error {
	// Insert data into the database
	result, err := db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", timestamp)
	if err != nil {
		log.Printf("Error inserting timestamp into database: %v\n", err)
		return fmt.Errorf("error inserting timestamp: %v", err)
	}

	// Log the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v\n", err)
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	log.Printf("Time log inserted successfully. Rows affected: %d\n", rowsAffected)
	return nil
}

// API Handler
func getCurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// timezone for Toronto
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		log.Printf("Error loading Toronto timezone: %v\n", err)
		http.Error(w, "Error loading timezone", http.StatusInternalServerError)
		return
	}

	// Get current time in Toronto
	currentTime := time.Now().In(location)

	// Push data into the database
	err = DatatoDatabase(currentTime)
	if err != nil {
		log.Printf("Error processing time log to database: %v\n", err)
		http.Error(w, "Sorry, processing of time log to the server failed!", http.StatusInternalServerError)
		return
	}

	// Prepare JSON response
	response := TimeResponse{CurrentTime: currentTime.Format("2006-01-02 15:04:05")}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Initialize database connection
	connDB()
	// Close database connection when main function exits
	defer db.Close()

	// API endpoint
	http.HandleFunc("/current-time", getCurrentTimeHandler)

	// Start the server on port 8080
	port := 8080
	log.Printf("Server running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
