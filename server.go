package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	r, err := json.Marshal("URL-shortening Service. Version 0.1.0")
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(r)
	if err != nil {
		log.Panic(err)
	}
}

func server() {
	host := getEnvValue("HOST")
	port := getEnvValue("PORT")
	log.Println("Server started!")
	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
