package Cache

import "github.com/go-redis/redis"

var (
	Cache = New()
)

func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "127.0.0.1",
		Password: "111222",
		DB: 0,
	})
}