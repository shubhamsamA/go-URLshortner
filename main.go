package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
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

func main() {
	fmt.Println("hello world")

	fmt.Println("Starting server on port 3000...")
	err := http.ListenAndServe(":3000",nil)
	if err != nil {
		fmt.Println("Error starting server:",err)
		
	}


}