package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	r, err := json.Marshal("URL-shortening Service. Version 1.0")
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(r)
	if err != nil {
		log.Panic(err)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	userURL := r.URL.Query().Get("url")
	userInput := r.URL.Query().Get("input")
	var genURL string
	u := fromDB(userURL, "userURL")
	if u == "URL not found" {
		genURL = getURL(userInput)
		if err := toDB(userURL, genURL); err != nil {
			log.Panic(err)
		}
	} else {
		genURL = u
	}
	w.Header().Set("Content-Type", "text")
	generatedURL, err := json.Marshal(genURL)
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(generatedURL)
	if err != nil {
		log.Panic(err)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	url = fromDB(url, "generatedURL")
	res, err := json.Marshal(url)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "text")
	_, err = w.Write(res)
	if err != nil {
		log.Panic(err)
	}
}

func server() {
	host := getEnvValue("HOST")
	port := getEnvValue("PORT")
	log.Println("Server started!")
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/generate", post)
	http.HandleFunc("/get", get)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
