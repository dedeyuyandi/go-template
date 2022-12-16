package redis

import (
	"fmt"

	"github.com/dedeyuyandi/go-template/repository"
	"github.com/go-redis/redis/v7"
)

// Redis Struct
type redisClient struct {
	client *redis.Client
}

type RedisConfiguration struct {
	RDSHost     string
	RDSPort     string
	RDSPassword string
	RDSDB       int
}

// NewClientRedis...
func NewClientRedis(conf repository.RedisConfiguration) (*redisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.RDSHost, conf.RDSPort),
		Password: conf.RDSPassword,
		DB:       conf.RDSDB,
	})

	ping, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("Redis Says: ", ping)

	return &redisClient{
		client: client,
	}, nil
}
