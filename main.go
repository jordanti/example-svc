package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Message struct {
	Message string `json:"msg"`
}

func main() {
	sport := os.Getenv("PORT")
	port, err := strconv.Atoi(sport)
	if err != nil {
		log.Fatalf("Invalid PORT value specified: %s. %s", sport, err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var msg = Message{
			Message: "It's alive!",
		}
		b, _ := json.Marshal(msg)
		w.Header().Add("Content-Type", "application/json")
		w.Write(b)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
