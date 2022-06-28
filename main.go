package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "...http request header detail...")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "v1.0.1")
	io.WriteString(w, os.Getenv("VERSION"))
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	clienIP := r.RemoteAddr
	url := r.RequestURI
	method := r.Method
	ua := r.UserAgent()
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unKnown"
	}
	log.Printf("%s %s %s %s %s\n", clienIP, url, method, ua, hostname)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/healthz", HealthHandler)
	http.HandleFunc("/log", LogHandler)
	http.HandleFunc("/version", VersionHandler)
	http.ListenAndServe(":5678", nil)
}
