package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type ProjetoResponse  struct {
	Nome string `json:"nome"`
	Horario string `json:"horario"`
}

func main(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/projeto-korp", handleProjetoKorp)
	http.ListenAndServe(":"+port, mux)
}

func handleProjetoKorp(
	w http.ResponseWriter, 
	r *http.Request,
){
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	response := ProjetoResponse{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}