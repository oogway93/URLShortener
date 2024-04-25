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

	tmpl.Execute(w, data)
	initialURL := r.FormValue("initial-url")
	shortURL := r.FormValue("short-url")
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()
	ctx := context.Background()
	if len(shortURL) != 0 {
		convertedURL, err := urlShortener.ConvertedURL(ctx, rdb, shortURL)
		if err != nil {
			log.Println(err)
			return
		} else {
			jsonResponse, _ := json.Marshal(map[string]string{"Converted URL": convertedURL})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(jsonResponse)
		}
	}
	if len(initialURL) != 0 {
		shortenURL, err := urlShortener.ShortenURL(ctx, rdb, initialURL)
		if err != nil {
			log.Println(err)
			return
		} else {
			jsonResponse, _ := json.Marshal(map[string]string{"Shorten URL": shortenURL})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(jsonResponse)
		}
	}
}
