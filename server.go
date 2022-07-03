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

func get(w http.ResponseWriter, r *http.Request) {
	//userURL := req.URL.Query().Get("url")
	w.Header().Set("Content-Type", "text")
	genURL, err := json.Marshal(getURL())
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(genURL)
	if err != nil {
		log.Panic(err)
	}
}
func server() {
	host := getEnvValue("HOST")
	port := getEnvValue("PORT")
	log.Println("Server started!")
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", get)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
