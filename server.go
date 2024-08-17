package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config-map", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/healthz", HealthCheck)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hi! I'm %s and i'm %s y/o", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(".env")
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
	}
	fmt.Fprintf(w, "data: %v", data)
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	pw := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "Hi! I'm %s and pw is %s", user, pw)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
