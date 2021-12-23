package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var workCount = 0

func main() {
	startBackgroundWork()
	startMetricsServer()
}

func startBackgroundWork() {
	go func() {
		for {
			workCount++
			fmt.Printf("Doing important background work round %d\n", workCount)
			time.Sleep(5 * time.Second)
		}
	}()
}

func startMetricsServer() {
	log.Print("starting server...")
	http.HandleFunc("/", statusHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{} {"healthy": true, "workCount": workCount}
	jsonStatus, err := json.Marshal(status)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = w.Write(jsonStatus); err != nil {
		log.Fatal(err)
	}
}
