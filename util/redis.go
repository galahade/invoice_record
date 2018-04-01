package util

import (
	"fmt"
	"log"
	"sync"
	"github.com/garyburd/redigo/redis"
	"github.com/olebedev/config"
)

var redisPool *redis.Pool
var once sync.Once

// GetRedisPool return a redis pool to connect to redis server
func GetRedisPool(cfg config.Config) *redis.Pool {
	once.Do(func() {
		tmp, err := cfg.String("redis.url")
		pwd, err1 := cfg.String("redis.password")
		if err != nil || err1 != nil {
			panic(fmt.Sprintf("Fail to read redis server config"))
		}

		redisPool = redis.NewPool(func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", tmp, redis.DialPassword(pwd))

			if err != nil {
				log.Fatalf("connect to %s failed with err : %s. ", tmp, err)
				return nil, err
			}
			return c, err
		}, 10)
	})
	return redisPool
}
