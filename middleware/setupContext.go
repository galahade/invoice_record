package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/olebedev/config"
	"github.com/garyburd/redigo/redis"
)

const(
	RedisConnKey = "redisConn"
	ProjectConfigKey = "cfg"
)

func SetupConfig(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ProjectConfigKey, cfg)
		c.Next()
	}
}

func SetupRedisConn(pool *redis.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(RedisConnKey, pool.Get())
		c.Next()
	}
}