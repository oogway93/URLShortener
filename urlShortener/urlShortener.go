package urlShortener

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

const KeyString = "urlshorten: "

func ShortenURL(ctx context.Context, rdb *redis.Client, url string) (string, error) {
	hash := sha1.New()
	hash.Write([]byte(url))
	shortURL := base64.URLEncoding.EncodeToString(hash.Sum(nil))[:8]

	if err := rdb.Set(ctx, KeyString+shortURL, url, 20*time.Second).Err(); err != nil {
		log.Println(err)
		return "", err
	}
	return shortURL, nil
}

func ConvertedURL(ctx context.Context, rdb *redis.Client, shortURL string) (string, error) {
	url, err := rdb.Get(ctx, KeyString+shortURL).Result()
	if err != nil {
		return "", err
	}
	return url, nil
}
