package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type message struct {
	X   int `json:"x"`
	Y   int `json:"y"`
	Sum int `json:"sum"`
}

func main() {
	sport := os.Getenv("PORT")
	port, err := strconv.Atoi(sport)
	if err != nil {
		log.Fatalf("Invalid PORT value specified: %s. %s", sport, err.Error())
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	sx := strings.TrimSpace(r.URL.Query().Get("x"))
	sy := strings.TrimSpace(r.URL.Query().Get("y"))

	if len(sx) == 0 || len(sy) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You must provide x and y values"))
		return
	}

	x, err := strconv.Atoi(sx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("x value is not an int"))
		return
	}

	y, err := strconv.Atoi(sy)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("y value is not an int"))
		return
	}

	sum := sum(x, y)
	var msg = message{
		X:   x,
		Y:   y,
		Sum: sum,
	}
	b, _ := json.Marshal(msg)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func sum(x int, y int) int {
	return x + y
}
