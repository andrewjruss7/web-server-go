package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World!")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the API endpoint")
}

func POSTRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metada metaData
	err := decoder.Decode(&metada)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error: %v\n", err)
		return
	}
	_, _ = fmt.Fprintf(w, "Payload %v\n", metada)
}

func UserPOSTRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error: %v\n", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}