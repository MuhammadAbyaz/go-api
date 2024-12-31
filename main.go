package main

import (
	"museum/api"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RUNNING"))
}
func main() {
	server := http.NewServeMux()
	server.HandleFunc("/health-check", healthCheck)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
	http.ListenAndServe(":8000", server)
}
