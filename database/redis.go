package database

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb redis.Client
