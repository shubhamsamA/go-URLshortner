package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type URL struct {
	ID string `json:"id"`
	OriginalURL string `json:"orgingal_url"`
	ShortURL string  `json:"short_url"`
	CreatedDate time.Time `json:"creted_date"`

}

var urlDB = make(map[string]URL)

func generateShortURL(originalURL string) string{
	hasher :=  md5.New()
	hasher.Write([]byte(originalURL))
	data :=hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

func createURL(originalURL string) string{
	shortURL := generateShortURL(originalURL)
	id :=shortURL
	urlDB[id] = URL{
		ID: id,
		OriginalURL: originalURL,
		ShortURL: shortURL,
		CreatedDate: time.Now(),
	}

	return shortURL
}

func getURL(id string) (URL, error){

	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}

	return url, nil

}

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello, world!")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request){
	var data struct{
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w,"Invalid request boady",http.StatusBadRequest)
		return
	}
	shortURL_:= createURL(data.URL)
	response := struct{
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(response)

}

func redirectURLHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w,"URL not foumd",http.StatusBadRequest)
	}
	http.Redirect(w, r, url.OriginalURL,http.StatusFound)
}
func main() {
	fmt.Println("hello world")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/shorten",ShortURLHandler)
	http.HandleFunc("/redirect/",redirectURLHandler)

	fmt.Println("Starting server on port 3000...")
	port := os.Getenv("PORT")
if port == "" {
    port = "3000" // fallback for local testing
}
http.ListenAndServe(":"+port, nil)

	}
