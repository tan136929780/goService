package redis

import (
	"github.com/go-redis/redis"
	"goMicroService/server-service/utils/config"
	"runtime"
	"time"
)

const (
	poolSizeFactor = 10
)

var (
	Client *redis.Client
)

func Setup() {
	if Client == nil {
		Client = newClient()
	}
}

func GetClient() *redis.Client {
	return Client
}

func newClient() *redis.Client {
	PoolSize := config.GetInt("db.redis.PoolSize")
	if PoolSize == 0 {
		PoolSize = poolSizeFactor
	}
	return redis.NewClient(&redis.Options{
		Addr:         config.GetString("db.redis.Host"),
		DialTimeout:  300 * time.Millisecond,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
		PoolSize:     PoolSize * runtime.NumCPU(),
		PoolTimeout:  3 * time.Second,
	})
}
