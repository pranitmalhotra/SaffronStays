package main

import (
    "fmt"
    "log"
    "net/http"
    "saffronstays-api/handlers"
    "saffronstays-api/config"
    "github.com/gorilla/mux"
)

// main initializes the application by setting up the database connection, 
// configuring the HTTP router, and starting the web server.
func main() {
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    router := mux.NewRouter()

    router.HandleFunc("/api/room/{room_id}", handlers.RoomHandler(db)).Methods("GET")

    fmt.Println("Server listening on port 8000...")
    log.Fatal(http.ListenAndServe(":8000", router))
}
