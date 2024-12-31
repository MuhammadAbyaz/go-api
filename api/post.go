package api

import (
	"encoding/json"
	"museum/data"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
	var newExhibition data.Exhibition
	err := json.NewDecoder(r.Body).Decode(&newExhibition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	data.PostExhibition(newExhibition)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}
