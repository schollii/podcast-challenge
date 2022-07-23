package main

import (
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var httpMux = mux.NewRouter()
var podcastExists = true
var RAND_ID int

const podcastName = "my-file.mp3"

func init() {
	rand.Seed(time.Now().UnixNano())
	RAND_ID = rand.Intn(100000)
}

func isRequestOnKnownPodcast(r *http.Request) bool {
	return mux.Vars(r)["name"] == podcastName
}

func Health(w http.ResponseWriter, _ *http.Request) {
	log.Printf("Service ID %d handling health check", RAND_ID)
	w.WriteHeader(http.StatusOK)
}

func Put(w http.ResponseWriter, r *http.Request) {
	log.Printf("Service ID %d handling podcast modification", RAND_ID)

	if !isRequestOnKnownPodcast(r) {
		http.Error(w, "unknown file", http.StatusBadRequest)
		return
	}

	if podcastExists {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "unknown file", http.StatusNotFound)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	log.Printf("Service ID %d handling podcast fetching", RAND_ID)

	if !(isRequestOnKnownPodcast(r) && podcastExists) {
		http.Error(w, "podcast not found", http.StatusNotFound)
		return
	}

	_, err := w.Write([]byte("podcast!\n"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("Service ID %d handling podcast deletion", RAND_ID)

	if !isRequestOnKnownPodcast(r) {
		http.Error(w, "podcast not found", http.StatusNotFound)
		return
	}

	if podcastExists {
		podcastExists = false
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "podcast not found", http.StatusNotFound)
	}
}

func main() {
	log.Printf(`Service UID %d`, RAND_ID)

	httpMux.HandleFunc("/v1/podcast/{name}", Put).Methods("PUT")
	httpMux.HandleFunc("/v1/podcast/{name}", Get).Methods("GET")
	httpMux.HandleFunc("/v1/podcast/{name}", Delete).Methods("DELETE")
	httpMux.HandleFunc("/health", Health).Methods("HEAD")

	log.Print("Starting server")
	log.Fatal(http.ListenAndServe(":8080", httpMux))
}
