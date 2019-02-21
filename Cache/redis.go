package Cache

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
)

type RedisStorage struct {
}

var (
	instance *RedisStorage
	once sync.Once
	client *redis.Client
	err error
)

func Instance() *RedisStorage {
	once.Do(func() {
		instance = &RedisStorage{}
	})
	return instance
}

func (r *RedisStorage) Init() (issue error) {
	var (
		addr, password string
		db int
	)
	addr = "127.0.0.1:6379"
	password = "111222"
	db = 1
	client = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: db,
	})

	_, err = client.Ping().Result()

	if err != nil {
		return err
	}

	return nil
}

func (r *RedisStorage) Set(key, value string, expiration time.Duration) error {
	err := client.Set(key, value, expiration).Err()
	return err
}

func (r *RedisStorage) Get(key string) (string, error) {
	result, err := client.Get(key).Result()
	return result, err
}

func (r *RedisStorage) Del(key string) error {
	err := client.Del(key).Err()
	return err
}