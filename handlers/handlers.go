package handlers

import (
	"URLShortener/urlShortener"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Title string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		panic(err)
	}
	data := &ViewData{
		Title: "Url shortener",
	}
	inputURL := r.FormValue("input-url")
	tmpl.Execute(w, data)

	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()
	ctx := context.Background()
	shortURL, err := urlShortener.ShortenURL(ctx, rdb, inputURL)
	if err != nil {
		log.Println(err)
		return
	} else {
		jsonResponse, _ := json.Marshal(map[string]string{"short-url": shortURL})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}

func RedirectToURL(rdb *redis.Client, shortURL string) (string, error) {
	ctx := context.Background()
	url, err := rdb.Get(ctx, urlShortener.KeyString+shortURL).Result()
	if err != nil {
		return "", err
	}

	return url, nil
}
