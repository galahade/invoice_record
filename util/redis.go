package util

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	"log"
)

var redisClient *redis.Pool
var once sync.Once
var redisServerString = "192.168.200.10:6379"

func GetRedisClient() *redis.Pool {
	once.Do(func() {
		redisClient = redis.NewPool(func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServerString)

			if err != nil {
				log.Fatalf("connect to %s failed. ",redisServerString)
				return nil, err
			}
			return c, err
		}, 10)
	})
	return redisClient
}
