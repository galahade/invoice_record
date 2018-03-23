package util

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	"log"
	"fmt"
	"github.com/olebedev/config"
)

var redisClient *redis.Pool
var once sync.Once

func GetRedisClient(cfg config.Config) *redis.Pool {
	once.Do(func() {
		tmp, err := cfg.String("redis.url")
		if err != nil {
			panic(fmt.Sprintf("Fail to read redis server config"))
		}

		redisClient = redis.NewPool(func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", tmp)

			if err != nil {
				log.Fatalf("connect to %s failed with err : %s. ",tmp, err)
				return nil, err
			}
			return c, err
		}, 10)
	})
	return redisClient
}
